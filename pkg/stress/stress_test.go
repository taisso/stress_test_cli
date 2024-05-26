package stress_test

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taisso/stress-test/pkg/stress"
)

func TestStressTest(t *testing.T) {
	var count int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&count, 1)
		if count > 10 && count <= 20 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if count > 20 && count <= 50 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	nRequest := 100
	nConcurrency := 10
	s := stress.NewStress(server.URL, nRequest, nConcurrency)

	report := s.Run()

	assert.Equal(t, nRequest, report.TotalRequests)
	assert.Equal(t, 60, report.SuccessRequests)

	assert.Equal(t, 60, report.StatusCounts[http.StatusOK])
	assert.Equal(t, 10, report.StatusCounts[http.StatusBadRequest])
	assert.Equal(t, 30, report.StatusCounts[http.StatusNotFound])
}

func BenchmarkStressTestStress(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	nRequest := 10000
	nConcurrency := 50
	s := stress.NewStress(server.URL, nRequest, nConcurrency)
	for range b.N {
		s.Run()
	}
}

package stress

import (
	"net/http"
	"sync"
	"time"
)

type Stress struct {
	url       string
	nRequest  int
	nCurrency int
}

type stressReport struct {
	TotalTime       time.Duration
	TotalRequests   int
	SuccessRequests int
	StatusCounts    map[int]int
}

func NewStress(url string, nRequest int, nCurrency int) *Stress {
	return &Stress{url, nRequest, nCurrency}
}

func (s *Stress) Run() *stressReport {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	start := time.Now()

	statusCounts := make(map[int]int)
	var totalRequests, successRequests = 0, 0

	requestsPerGoroutine := s.nRequest / s.nCurrency
	remainingRequests := s.nRequest % s.nCurrency

	for index := range s.nCurrency {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			requestsToExecute := requestsPerGoroutine
			if index < remainingRequests {
				requestsToExecute++
			}
			for range requestsToExecute {
				resp, err := http.Get(s.url)
				if err != nil {
					return
				}
				defer resp.Body.Close()

				mutex.Lock()
				statusCounts[resp.StatusCode]++
				if resp.StatusCode == http.StatusOK {
					successRequests++
				}
				totalRequests++
				mutex.Unlock()
			}
		}(index)
	}
	wg.Wait()

	elapsed := time.Since(start)

	return &stressReport{
		TotalTime:       elapsed,
		TotalRequests:   totalRequests,
		SuccessRequests: successRequests,
		StatusCounts:    statusCounts,
	}
}

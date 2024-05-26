package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/taisso/stress-test/pkg/stress"
)

func main() {
	url := flag.String("url", "", "URL of the service to be tested")
	nRequest := flag.Int("requests", 0, "Total number of requests")
	nConcurrency := flag.Int("concurrency", 0, "Number of simultaneous calls")

	flag.Parse()

	if *url == "" {
		log.Fatal("URL not provided")
	}

	if *nRequest == 0 {
		log.Fatal("requests must be greater than zero")
	}

	if *nConcurrency == 0 {
		log.Fatal("concurrency must be greater than zero")
	}

	s := stress.NewStress(*url, *nRequest, *nConcurrency)

	report := s.Run()

	fmt.Printf("Total time: %s\n", report.TotalTime)
	fmt.Printf("Total Requests: %d\n", report.TotalRequests)
	fmt.Printf("Successful Requests (HTTP 200): %d\n", report.SuccessRequests)
	fmt.Println("Status Code Distribution:")
	for status, count := range report.StatusCounts {
		fmt.Printf("\to status %d ocorreu %d vezes\n", status, count)
	}
}

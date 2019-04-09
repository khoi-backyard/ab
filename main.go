package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type responseMeta struct {
	statusCode    int
	contentLength int64
	responseTime  time.Duration
}

func main() {
	requestNums := flag.Int("n", 1, "number of requests")
	concurrency := flag.Int("c", 1, "number of requests to perform at once")

	flag.Parse()

	if flag.NArg() == 0 || *requestNums <= 0 || *concurrency < 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	host := flag.Arg(0)
	fmt.Printf("🚀 Benchmarking %v (be patient)...\n", host)

	tasks := make(chan struct{})
	var wg sync.WaitGroup
	var result []*responseMeta
	var resultMux sync.Mutex

	for worker := 0; worker < *concurrency; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range tasks {
				fmt.Println(i)
				resultMux.Lock()
				result = append(result, get(host))
				resultMux.Unlock()
			}
		}()
	}

	for i := 0; i < *requestNums; i++ {
		tasks <- struct{}{}
	}

	close(tasks)
	wg.Wait()

	var totalTime time.Duration

	for _, v := range result {
		totalTime += v.responseTime
	}

	mean := int(totalTime/time.Millisecond) / len(result)

	fmt.Printf("Time per request: %v [ms] (mean)\n", mean)
}

func get(url string) *responseMeta {
	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	return &responseMeta{
		statusCode:    res.StatusCode,
		contentLength: res.ContentLength,
		responseTime:  time.Now().Sub(start),
	}
}

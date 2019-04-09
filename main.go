package main

import (
	"flag"
  "fmt"
	"os"
  "net/http"
  "time"
  "sync"
)

type responseMeta struct {
  statusCode int
  contentLength int64
  responseTime time.Duration
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

  tasks := make(chan struct{})
  var wg sync.WaitGroup

  for worker := 0; worker < *concurrency; worker++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for i := range tasks {
        meta := get(host)
        fmt.Printf("%v %v+\n",i,  meta)
      }
    }()
  }

  for i:=0; i < *requestNums; i++ {
    tasks <- struct{}{}   
  }

  close(tasks)
  wg.Wait()
}

func get(url string) *responseMeta {
  start := time.Now()

  res, err := http.Get(url);

  if err != nil {
    panic(err)
  }

  return &responseMeta {
    statusCode: res.StatusCode,
    contentLength: res.ContentLength,
    responseTime: time.Now().Sub(start),
  }
}

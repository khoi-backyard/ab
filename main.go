package main

import (
	"flag"
	"fmt"
	"os"
  "net/http"
  "time"
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

	fmt.Printf("reqNums %v \n", *requestNums)
	fmt.Printf("concurrency %v \n", *concurrency)
	fmt.Printf("ddosing %v \n", host)

  responseMeta := get("http://httpun.org/ip")

  fmt.Printf("%+v \n", responseMeta)
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

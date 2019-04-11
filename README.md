# CoderSchool Golang Course Submission

2. **Time spent: 2 hours**

## Usage

```bash
ab -n 500 -c 10 https://localhost:8080
```

## Set of User Stories

### Required
* [x] Command-line argument parsing
* [x] Input params
   * [x] Requests - Number of requests to perform
   * [x] Concurrency - Number of multiple requests to make at a time
   * [x] URL - The URL for testing.
* [x] Prints use information if wrong arguments provided
* [x] Implements  HTTP load and summarize it
* [x] Concurrency must be implemented with goroutine.


### Bonus
* [ ] Extend input params with: 
   * [ ] Timeout - Seconds to max. wait for each response
   * [ ] Timelimit - Maximum number of seconds to spend for benchmarking
* [ ] Prints key metrics of summary, such:
   * [ ] Server Hostname
   * [ ] Server Port
   * [ ] Document Path
   * [ ] Document Length
   * [ ] Concurrency Level
   * [ ] Time taken for tests
   * [x] Complete requests
   * [x] Failed requests
   * [ ] Total transferred
   * [ ] Requests per second
   * [ ] Time per request
   * [ ] Time per request
   * [ ] Transfer rate

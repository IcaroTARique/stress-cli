package api

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Request struct {
	Url     string
	Jobs    int
	Workers int
}

type Response struct {
	TotalTime  time.Duration
	ReqAmmount int
	Responses  map[string]int
}

func NewRequest() Request {
	return Request{}
}

func (r *Request) SetUrl(url string) {
	r.Url = url
}

func (r *Request) SetJobs(jobs int) {
	r.Jobs = jobs
}

func (r *Request) SetWorkers(workers int) {
	r.Workers = workers
}

func (r *Request) GoRequest(verbose bool) Response {
	initialTime := time.Now()
	responseCodes := map[string]int{}

	wg := sync.WaitGroup{}
	var respWg sync.WaitGroup
	defer wg.Wait()

	readyJobs := make(chan int, r.Jobs)
	jobsResponse := make(chan int, r.Jobs)

	for i := 0; i < r.Workers; i++ {
		wg.Add(1)
		go r.GetRequest(&wg, readyJobs, jobsResponse, verbose)
	}

	respWg.Add(1)
	go func() {
		defer respWg.Done()
		for t := range jobsResponse {
			if verbose {
				fmt.Println("Response code: ", t)
			}
			responseCodes[strconv.Itoa(t)]++
		}
	}()

	for i := 0; i < r.Jobs; i++ {
		readyJobs <- i
	}

	close(readyJobs)
	wg.Wait()
	close(jobsResponse)
	respWg.Wait()

	endTime := time.Now()

	return Response{
		TotalTime:  endTime.Sub(initialTime),
		ReqAmmount: r.Jobs,
		Responses:  responseCodes,
	}

}

func (r *Request) GetRequest(wg *sync.WaitGroup, ch1 <-chan int, ch2 chan<- int, verbose bool) {
	defer wg.Done()
	for x := range ch1 {
		res, _ := http.Get(r.Url)
		res.Body.Close()
		ch2 <- res.StatusCode
		if verbose {
			fmt.Println(x, " message is status code: ", res.StatusCode)
		}

	}
}

package main

import (
	"flag"
	"fmt"
	"net/http"
	"slices"
	"time"
)

type Task struct {
	Url    string
	Method string
}

type TaskResult struct {
	StatusCode  int
	Error       bool
	RequestTime float64
}

type Worker struct {
}

func (w *Worker) Start(tasks chan Task, results chan TaskResult) {
	for task := range tasks {
		var result TaskResult
		request, err := http.NewRequest(task.Method, task.Url, nil)
		if err != nil {
			panic(err)
		}
		start := time.Now()
		response, err := http.DefaultClient.Do(request)
		end := time.Now()
		result.RequestTime = end.Sub(start).Seconds()
		if err != nil {
			result.Error = true
		} else {
			result.StatusCode = response.StatusCode
		}
		results <- result
	}
}

func mean(s []float64) float64 {
	var sum float64
	for _, x := range s {
		sum += x
	}
	return sum / float64(len(s))
}

func main() {
	url := flag.String("u", "", "URL")
	method := flag.String("method", "", "HTTP Method")
	expect := flag.Int("expect", 200, "Expected status code")
	n := flag.Int("n", 10, "Number of requests")
	c := flag.Int("c", 10, "Number of concurrect requests")

	flag.Parse()

	tasks := make(chan Task, *n)
	results := make(chan TaskResult, *n)
	for i := 0; i < *c; i++ {
		worker := Worker{}
		go func() {
			worker.Start(tasks, results)
		}()
	}

	for i := 0; i < *n; i++ {
		task := Task{Url: *url, Method: *method}
		tasks <- task
	}
	close(tasks)

	failures := 0
	requestTimes := make([]float64, 0)
	for i := 0; i < *n; i++ {
		result := <-results
		if result.Error || result.StatusCode != *expect {
			failures += 1
		}
		requestTimes = append(requestTimes, result.RequestTime)
	}
	close(results)

	fmt.Printf("Successes\t%d\n", *n-failures)
	fmt.Printf("Failures (5xx)\t%d\n", failures)
	fmt.Printf("Request Time (s) (Min, Max, Mean): %.3f, %.3f, %.3f\n", slices.Min(requestTimes), slices.Max(requestTimes), mean(requestTimes))
}

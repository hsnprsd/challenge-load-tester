package main

import (
	"flag"
	"fmt"
	"net/http"
)

type Task struct {
    Url string
}

type TaskResult struct {
    StatusCode int
    Error bool
}

type Worker struct {
}

func (w *Worker) Start(tasks chan Task, results chan TaskResult) {
    for task := range tasks {
        var result TaskResult
        response, err := http.Get(task.Url)
        if err != nil {
            result.Error = true
        } else {
            result.StatusCode = response.StatusCode
        }
        results <- result
    }
}

func main() {
    url := flag.String("u", "", "URL")
    n := flag.Int("n", 10, "Number of requests to make")
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
        task := Task{Url: *url}
        tasks <- task
    }
    close(tasks)
    failures := 0
    for i := 0; i < *n; i++ {
        result := <-results
        if result.Error || (500 <= result.StatusCode && result.StatusCode < 600) {
            failures += 1
        }
    }
    close(results)
    fmt.Printf("Successes: %d\n", *n - failures)
    fmt.Printf("Failures: %d\n", failures)
}

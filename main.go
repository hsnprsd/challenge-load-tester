package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
    url := flag.String("u", "", "URL")
    n := flag.Int("n", 10, "Number of requests to make")

    flag.Parse()

    for i := 0; i < *n; i++ {
        res, err := http.Get(*url)
        if err != nil {
            panic(err)
        }
        fmt.Printf("Response code: %d\n", res.StatusCode)
    }
}

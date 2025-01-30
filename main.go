package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
    url := os.Args[1]
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Response code: %d\n", res.StatusCode)
}

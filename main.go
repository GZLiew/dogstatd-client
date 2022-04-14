package main

import (
    "log"
    "github.com/DataDog/datadog-go/v5/statsd"
)

func main() {
    statsd, err := statsd.New("127.0.0.1:8125")
    if err != nil {
        log.Fatal(err)
    }
}


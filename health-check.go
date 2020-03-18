package main

import (
	"flag"
	"net/http"
	"os"
	"time"
)

var timeout time.Duration

func main() {
	flag.DurationVar(&timeout, "t", 3*time.Second, "HTTP connection timeout in seconds")
	flag.Parse()

	c := &http.Client{
		Timeout: timeout,
	}

	_, err := c.Get(flag.Arg(0))
	if err != nil {
		os.Exit(1)
	}
}

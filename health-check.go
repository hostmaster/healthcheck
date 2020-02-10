package main

import (
	"net/http"
	"os"
)

func main() {
	_, err := http.Get(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
}

package main

import (
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8080/healthcheck")

	if err != nil || resp.StatusCode != 200 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

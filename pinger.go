package main

import (
	"net/http"
	"time"
)

// HTTP Get ping
func getPing(url string) PingResponse {
	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Get(url)

	// HTTP GET Error
	if err != nil {
		return PingResponse{url,false, err.Error(), 0 }
	}

	return PingResponse{url,true, "", response.StatusCode}
}

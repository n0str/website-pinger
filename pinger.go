package main

import (
	"net/http"
	"time"
)

// HTTP Get ping
func getPing(url string) PingResponse {
	//time.Sleep(1 * time.Second)
	timeout := time.Duration(5 * time.Second)
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

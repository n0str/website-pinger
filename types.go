package main

type PingResponse struct {
	url string
	result bool
	message string
	statusCode int
}

type CheckRule struct {
	url string
	desiredStatusCode int
}

type job struct {
	task     CheckRule
}

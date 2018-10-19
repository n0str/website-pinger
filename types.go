package main

type PingResponse struct {
	url string
	result bool
	message string
	statusCode int
}

type CheckRule struct {
	Url string
	DesiredStatusCode int
}

type job struct {
	task     CheckRule
}

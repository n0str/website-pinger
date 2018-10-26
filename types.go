package main

// PingResponse – HTTP Response from Website
type PingResponse struct {
	url        string
	result     bool
	message    string
	statusCode int
}

// CheckRule – Ping settings for the URL
type CheckRule struct {
	URL               string
	DesiredStatusCode int
	InformerPayload   InformerData
}

// InformerData – Payload data for Informer
type InformerData struct {
	Type    int
	Payload string
}

// Running ping Job
type job struct {
	task CheckRule
}

const (
	codexBotInformer = iota
	hawkInformer     = iota
)

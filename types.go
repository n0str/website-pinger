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
	InformerPayload InformerData
}

type InformerData struct {
	Type int
	Payload string
}

type job struct {
	task     CheckRule
}

const (
	codexBotInformer = iota
	hawkInformer = iota
)
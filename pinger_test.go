package main

import (
	"testing"
)

func TestGetPingSuccess(t *testing.T) {
	response := getPing("https://ifmo.su")
	if !response.result {
		t.Errorf("Response should be true, but it is %t", response.result)
		t.Errorf("Ping get failed: code=%d, message=%s.", response.statusCode, response.message)
	}
}

func TestGetPingDNSFail(t *testing.T) {
	response := getPing("https://undefined.undefined")
	if response.result {
		t.Errorf("Response should be false, but it is %t", response.result)
		t.Errorf("Ping get success: code=%d, message=%s.", response.statusCode, response.message)
	}
	if response.statusCode != 0 {
		t.Errorf("Response statusCode should be 0, but it is %d", response.statusCode)
	}
}

func TestGetPing404Error(t *testing.T) {
	response := getPing("https://hawk.so/123")
	if response.statusCode != 404 {
		t.Errorf("Response statusCode should be 404, but it is %d", response.statusCode)
	}
}

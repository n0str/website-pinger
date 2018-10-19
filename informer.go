package main

import "fmt"

func inform(rule CheckRule, response PingResponse) {
	fmt.Printf("Inform about %s\n", response.message)
}
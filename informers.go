package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func inform(rule CheckRule, response PingResponse) {
	var message string
	fmt.Printf("Inform %s", rule.Url)
	if !response.result {
		message = fmt.Sprintf("Failed to connect: %s.\n_%s_", rule.Url, response.message)
	} else {
		message = fmt.Sprintf("Status code mismatched: %s.\n_%d_", rule.Url, response.statusCode)
	}
	if rule.InformerPayload.Type == codexBotInformer {
		codexBotInform(rule.InformerPayload.Payload, message)
	}
}

func codexBotInform(payload string, message string)  {
	botUrl := "https://notify.bot.ifmo.su/u/" + payload
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	_, err := client.PostForm(botUrl, url.Values{"message": {message}, "disable_web_page_preview": {"true"}, "parse_mode": {"Markdown"}})

	// HTTP POST Error
	if err != nil {
		log.Printf("[inform] codexBotInform Exception: %s", err.Error())
		return
	}
}
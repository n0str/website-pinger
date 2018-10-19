package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var rules []CheckRule
var rulesMap = make(map[string]CheckRule)
var maxWorkers = 100
var maxQueueSize = 5
var jobs chan job

func main() {
	maxQueueSize = *flag.Int("max_queue_size", 100, "The size of job queue")
	maxWorkers   = *flag.Int("max_workers", 5, "The number of workers to start")
	var (
		port         = flag.String("port", "8080", "The server port")
	)
	flag.Parse()

	jobs = make(chan job, maxQueueSize)

	println("Hello World!")
	rand.Seed(time.Now().UTC().UnixNano())
	loadRules()
	initAPIHandlers()
	runLoop()
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func loadRules() {
	var newRules []CheckRule
	newRules = append(newRules, CheckRule{"https://github.com", 200} )
	for i := 0; i <= 5; i++ {
		newUrl := fmt.Sprintf("https://github.com/%d", i)
		newRule := CheckRule{newUrl, 200}
		rulesMap[newUrl] = newRule
		newRules = append(newRules, newRule)
	}
	rules = newRules
}

func doTask(rule CheckRule) {
	r := getPing(rule.url)
	if r.result {
		fmt.Printf("Url %s [OK]\n", rule.url)
	} else {
		fmt.Printf("Url %s [FAIL] - status=%d, cause=%s\n", rule.url, r.statusCode, r.message)
	}
}

// Endless loop with 5-second time ticker
func runLoop() {
	var endWaiter sync.WaitGroup
	endWaiter.Add(1)
	ticker := time.NewTicker(10 * time.Second)

	var signalChannel chan os.Signal
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	for i := 1; i <= maxWorkers; i++ {
		go func(i int) {
			for j := range jobs {
				doTask(j.task)
			}
		}(i)
	}

	go func() {
		runTasks()
		for {
			select {
			case <- ticker.C:
				runTasks()
			case <- signalChannel:
				ticker.Stop()
			}
		}
	}()
}

func runTasks() {
	//for _, j  := range rules {
	//	jobs <- job{j}
	//}
	for _, value := range rulesMap {
		jobs <- job{value}
	}
}

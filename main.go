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

var rulesMap = make(map[string]CheckRule)
var maxWorkers = 100
var maxQueueSize = 5
var jobs chan job

func main() {
	SetupLogs()

	maxQueueSize = *flag.Int("max_queue_size", 100, "The size of job queue")
	maxWorkers = *flag.Int("max_workers", 5, "The number of workers to start")
	var (
		port = flag.String("port", "8080", "The server port")
	)
	flag.Parse()

	jobs = make(chan job, maxQueueSize)

	println("Hello World!")
	rand.Seed(time.Now().UTC().UnixNano())

	dbReload()
	initAPIHandlers()
	runLoop()
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func doTask(rule CheckRule) {
	r := getPing(rule.URL)
	if r.result && r.statusCode == rule.DesiredStatusCode {
		fmt.Printf("Url %s [OK]\n", rule.URL)
	} else {
		// Inform about failure
		inform(rule, r)
	}
}

// Endless loop with 5-second time ticker
func runLoop() {
	var endWaiter sync.WaitGroup
	endWaiter.Add(1)
	ticker := time.NewTicker(15 * time.Second)

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
			case <-ticker.C:
				runTasks()
			case <-signalChannel:
				ticker.Stop()
			}
		}
	}()
}

func runTasks() {
	for _, value := range rulesMap {
		jobs <- job{value}
	}
}

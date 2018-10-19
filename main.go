package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

var rules []CheckRule
var maxQueueSize = 100
var maxWorkers = 5
var jobs = make(chan job, maxQueueSize)

func main() {
	println("Hello World!")
	rand.Seed(time.Now().UTC().UnixNano())
	loadRules()
	runLoop()
}

func loadRules() {
	var newRules []CheckRule
	newRules = append(newRules, CheckRule{"https://github.com", 200} )
	for i := 0; i <= 10; i++ {
		newRule := CheckRule{fmt.Sprintf("https://github.com/%d", i), 200}
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
	ticker := time.NewTicker(1 * time.Second)

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
				endWaiter.Done()
			}
		}
	}()
	endWaiter.Wait()
}

func runTasks() {
	for _, j  := range rules {
		jobs <- job{j}
	}
}

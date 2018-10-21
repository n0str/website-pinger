package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func dbSet(hostname string, rule CheckRule) {
	filename := fmt.Sprintf("%s%s__%s.db", dbPath, hostname, GetMD5Hash(rule.Url))
	Save(filename, rule)
}

func dbGet(filename string, rule *CheckRule) bool {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("[database] dbReload() Open Exception: %p", err.Error())
		return false
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Printf("[database] dbReload() Read Exception: %p", err.Error())
		return false
	}
	json.Unmarshal(data[:count], rule)
	return true
}

func dbReload() {
	files, err := ioutil.ReadDir(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".db") {
			var newRule CheckRule
			if dbGet(dbPath + f.Name(), &newRule) != true {
				log.Printf("[database] dbReload() Cannot unmarshal file: %s", f.Name())
				continue
			}
			rulesMap[newRule.Url] = newRule
		}
	}
	log.Println("[database] dbReload() Success")
}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// dbSet – save rule to database
func dbSet(hostname string, rule CheckRule) {
	filename := fmt.Sprintf("%s%s__%s.db", dbPath, hostname, GetMD5Hash(rule.URL))
	Save(filename, rule)
}

func dbDelete(hostname string, ruleURL string) bool {
	filename := fmt.Sprintf("%s%s__%s.db", dbPath, hostname, GetMD5Hash(ruleURL))
	err := os.Remove(filename)
	if err != nil {
		log.Printf("[database] dbDelete() Remove Exception: %s", err.Error())
		return false
	}
	return true
}

func dbGet(filename string, rule *CheckRule) bool {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("[database] dbReload() Open Exception: %s", err.Error())
		return false
	}
	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		log.Printf("[database] dbReload() Read Exception: %s", err.Error())
		return false
	}
	err = json.Unmarshal(data[:count], rule)
	if err != nil {
		log.Printf("[database] dbReload() Unmarshall Exception: %s", err.Error())
		return false
	}
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
			if dbGet(dbPath+f.Name(), &newRule) != true {
				log.Printf("[database] dbReload() Cannot unmarshal file: %s", f.Name())
				continue
			}
			rulesMap[newRule.URL] = newRule
		}
	}
	log.Println("[database] dbReload() Success")
}

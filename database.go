package main

import (
	"fmt"
	"os"
)

func dbSet(hostname string, url string) {
	filename := fmt.Sprintf("./rules/%s__%s", hostname, GetMD5Hash(url))

	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		return
	}

	_, err = f.WriteString("writes\n")
	if err != nil {
		return
	}

	f.Sync()
}
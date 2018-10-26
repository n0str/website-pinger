package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func initAPIHandlers() {
	http.HandleFunc("/api/set", func(w http.ResponseWriter, r *http.Request) {
		apiSetHandler(w, r)
	})
	http.HandleFunc("/api/get", func(w http.ResponseWriter, r *http.Request) {
		apiGetHandler(w, r)
	})
	http.HandleFunc("/api/reload", func(w http.ResponseWriter, r *http.Request) {
		apiReloadHandler(w, r)
	})
	http.HandleFunc("/api/list", func(w http.ResponseWriter, r *http.Request) {
		apiListHandler(w, r)
	})
	http.HandleFunc("/api/delete", func(w http.ResponseWriter, r *http.Request) {
		apiDeleteHandler(w, r)
	})
}

func apiSetHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Set statusCode and validate value.
	status := r.FormValue("status_code")
	if status == "" {
		http.Error(w, "You must specify a status code.", http.StatusBadRequest)
		return
	}
	statusCode, err := strconv.Atoi(status)
	if err != nil {
		http.Error(w, "Invalid status code.", http.StatusBadRequest)
		return
	}

	// Set url and validate value.
	ruleURL := r.FormValue("url")
	if ruleURL == "" {
		http.Error(w, "You must specify an URL.", http.StatusBadRequest)
		return
	}

	urlStruct, err := url.ParseRequestURI(ruleURL)
	if err != nil {
		http.Error(w, "URL is invalid", http.StatusBadRequest)
		return
	}

	print()

	informerType := r.FormValue("informer_type")
	informerTypeCode, err := strconv.Atoi(informerType)
	if err != nil {
		http.Error(w, "You must specify an informer type.", http.StatusBadRequest)
		return
	}

	informerPayload := r.FormValue("informer_payload")
	if informerPayload == "" {
		http.Error(w, "You must specify an informer payload.", http.StatusBadRequest)
		return
	}

	rulesMap[ruleURL] = CheckRule{ruleURL, statusCode, InformerData{informerTypeCode, informerPayload}}
	dbSet(urlStruct.Host, rulesMap[ruleURL])

	w.WriteHeader(http.StatusCreated)
	return
}

func apiGetHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Set url and validate value.
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "You must specify an URL.", http.StatusBadRequest)
		return
	}

	// Find rule for the URL
	data, ok := rulesMap[url]
	if ok {
		w.WriteHeader(http.StatusOK)
		// Return Rule Data
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Fatal error. JSON exception", http.StatusBadRequest)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	return
}

func apiReloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rulesMap = make(map[string]CheckRule)
	dbReload()
	log.Println("[API] apiReloadHandler() — Reload rules")
}

func apiListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var urls = []string{}
	for urlKey := range rulesMap {
		urls = append(urls, urlKey)
	}

	w.WriteHeader(http.StatusOK)
	// Return Rules URL
	err := json.NewEncoder(w).Encode(urls)
	if err != nil {
		http.Error(w, "Fatal error. JSON exception", http.StatusBadRequest)
		return
	}

	log.Println("[API] apiListHandler() — List rules")
}

func apiDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Make sure we can only be called with an HTTP POST request
	if r.Method != "DELETE" {
		w.Header().Set("Allow", "DELETE")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Set url and validate value.
	ruleURL := r.FormValue("url")
	if ruleURL == "" {
		http.Error(w, "You must specify an URL.", http.StatusBadRequest)
		return
	}

	urlStruct, err := url.ParseRequestURI(ruleURL)
	if err != nil {
		http.Error(w, "URL is invalid", http.StatusBadRequest)
		return
	}

	delete(rulesMap, ruleURL)
	dbDelete(urlStruct.Host, ruleURL)

	w.WriteHeader(http.StatusOK)
	return
}

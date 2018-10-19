package main

import (
	"net/http"
	"strconv"
)

func initAPIHandlers() {
	http.HandleFunc("/api/set", func(w http.ResponseWriter, r *http.Request) {
		apiSetHandler(w, r)
	})
	http.HandleFunc("/api/get", func(w http.ResponseWriter, r *http.Request) {
		apiGetHandler(w, r)
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
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "You must specify an URL.", http.StatusBadRequest)
		return
	}

	rulesMap[url] = CheckRule{url, statusCode}

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
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "You must specify an URL.", http.StatusBadRequest)
		return
	}

	rulesMap[url] = CheckRule{url, statusCode}

	w.WriteHeader(http.StatusCreated)
	return
}
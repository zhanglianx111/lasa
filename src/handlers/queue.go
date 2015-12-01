package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func HandlerGetQueue(w http.ResponseWriter, r *http.Request) {
	aQueue := make(map[string]string)

	queue, err := JenkinsClient.GetQueue()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if queue == nil {
		aQueue["name"] = ""
	} else {
		aQueue["name"] = queue.Base
	}

	jsonData, err := json.Marshal(aQueue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandlerGetQueueUrl(w http.ResponseWriter, r *http.Request) {
	aUrl := make(map[string]string)

	url := JenkinsClient.GetQueueUrl()
	if url == "" {
		aUrl["url"] = ""
	} else {
		aUrl["url"] = url
	}

	jsonData, err := json.Marshal(aUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprintf(w, "get queue")
}
package handlers

import (
	"encoding/json"
	"net/http"
	//"fmt"
)

type Job struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`
}

func HandlerGetInfo(w http.ResponseWriter, r *http.Request) {
	jenkinsInfo := make(map[string]interface{})
	var jobs []Job
	var views []interface{}

	info, err := JenkinsClient.Info()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//fmt.Println(info.AssignedLabels, info.Description, info.Jobs)
	// get mode of jenkins server
	jenkinsInfo["mode"] = info.Mode

	// jobs in jenkins server
	for _, j := range info.Jobs {
		jobs = append(jobs, Job{j.Name, j.Url, j.Color})
	}
	jenkinsInfo["jobs"] = jobs

	// nodes in jenkins server
	jenkinsInfo["nodename"] = info.NodeName

	// get slave agent port
	jenkinsInfo["slaveagentport"] = info.SlaveAgentPort

	// views
	for _, v := range info.Views {
		view := map[string]interface{}{"name": v.Name, "url": v.URL}
		views = append(views, view)
	}
	jenkinsInfo["views"] = views

	// NumExecutors
	jenkinsInfo["numberexecutors"] = info.NumExecutors
	jsonData, err := json.Marshal(jenkinsInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

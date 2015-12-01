package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"utils"
)

func HandlerCreateJob(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	jobid := params.Get(":jobid")

	job_data := getFileAsString("config.xml")

	job, err := JenkinsClient.CreateJob(job_data, jobid)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	//creating job is ok
	log.Debugf("job base:%s, job raw:", job.Base, job.Raw)
	data := map[string]string{"name": job.Raw.Name,
		"description": job.Raw.Description,
		"displayName": job.Raw.DisplayName,
		"url":         job.Raw.URL}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandlerDeleteJob(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	jobid := params.Get(":jobid")

	log.Debugf("delete job:%s", jobid)
	_, err := JenkinsClient.DeleteJob(jobid)
	if err != nil {
		log.Errorf(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	jsonData, err := json.Marshal(map[string]string{"name": jobid})
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandlerDisableJob(w http.ResponseWriter, r *http.Request) {
	jobid := r.URL.Query().Get(":jobid")

	log.Debugf("disable job: %s", jobid)
	job, err := JenkinsClient.GetJob(jobid)
	if err != nil {
		log.Errorf(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}

	_, err = job.Disable()
	if err != nil {
		log.Errorf("disable job:%s failed(code:%s)", jobid, err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "true")
	return
}

func HandlerEnableJob(w http.ResponseWriter, r *http.Request) {
	jobid := r.URL.Query().Get(":jobid")
	log.Debugf("enable job:%s", jobid)
	job, err := JenkinsClient.GetJob(jobid)
	if err != nil {
		log.Errorf(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}

	_, err = job.Enable()
	if err != nil {
		log.Errorf(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "true")
	return
}

func HandlerRenameJob(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	oldName := params.Get(":oldjobid")
	newName := params.Get(":newjobid")
	if oldName == "" || newName == "" {
		fmt.Fprintf(w, "oldname or newname is empty")
		return
	}

	job := JenkinsClient.RenameJob(oldName, newName)
	if job == nil {
		fmt.Fprintf(w, "job rename faild")
	}

	fmt.Fprintf(w, job.Base)
	return
}

func HandlerGetJob(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	jobid := params.Get(":jobid")

	job, err := JenkinsClient.GetJob(jobid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jobData := utils.AnalysisJob(job.Raw)

	jsonData, err := json.Marshal(jobData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandlerGetAllJobs(w http.ResponseWriter, r *http.Request) {
	var jobsData []map[string]interface{}
	jobs, _ := JenkinsClient.GetAllJobs()
	for _, job := range jobs {
		fmt.Println(job)
		jobData := utils.AnalysisJob(job.Raw)
		jobsData = append(jobsData, jobData)
	}

	jsonData, err := json.Marshal(jobsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getFileAsString(path string) string {
	fmt.Println(path)
	buf, err := ioutil.ReadFile("/Users/zhanglianxiang/workspace/jenkins_api/src/handlers/_tests/" + path)
	if err != nil {
		panic(err)
	}

	return string(buf)
}

func HandlerGetAllBuildIds(w http.ResponseWriter, r *http.Request) {
	bs := []map[string]interface{}{}
	jobid := r.URL.Query().Get(":jobid")
	if jobid == "" {
		fmt.Fprintf(w, "jobid is empty")
		return
	}

	job, err1 := JenkinsClient.GetJob(jobid)
	if err1 != nil {
		fmt.Fprintf(w, err1.Error())
		return
	}

	builds, err2 := job.GetAllBuildIds()
	if err2 != nil {
		fmt.Fprintf(w, err2.Error())
		return
	}

	for _, b := range builds {
		bs = append(bs, map[string]interface{}{"number": b.Number, "url": b.URL})
	}
	jsonData, err3 := json.Marshal(bs)
	if err3 != nil {
		fmt.Fprintf(w, err3.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandlerJobRunning(w http.ResponseWriter, r *http.Request) {
	jobid := r.URL.Query().Get(":jobid")
	if jobid == "" {
		fmt.Fprintf(w, "jobid is empty")
		return
	}

	job, err := JenkinsClient.GetJob(jobid)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	b, err1 := job.IsRunning()
	if err1 != nil {
		fmt.Fprintf(w, err1.Error())
		return
	}
	fmt.Println("b:", b)
	if b {
		fmt.Fprintf(w, "true")
	} else {
		fmt.Fprintf(w, "false")
	}
}

func HandlerBuildJob(w http.ResponseWriter, r *http.Request) {
	jobid := r.URL.Query().Get(":jobid")

	job, err := JenkinsClient.GetJob(jobid)
	if err != nil {
		fmt.Fprintf(w, "get job:%s failed", jobid)
		return
	}
	enable, _ := job.IsEnabled()
	if !enable {
		log.Infof("job:%s is disable, can't build", jobid)
		fmt.Fprintf(w, "job:%s is disable, can't build", jobid)
		return
	}
	queue, _ := job.IsQueued()
	if queue {
		log.Infof("job:%s already in queue")
		fmt.Fprintf(w, "job:%s already in queue", jobid)
		return
	}

	b, err := JenkinsClient.BuildJob(jobid)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	if b {
		fmt.Fprintf(w, "true")
	} else {
		fmt.Fprintf(w, "false")
	}
}

func HandlerStopBuild(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	jobid := q.Get(":jobid")
	number := q.Get(":number")
	n, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Errorf("string to int64 faild")
		return
	}
	build, err := JenkinsClient.GetBuild(jobid, n)
	if err != nil {
		log.Errorf("get build job:%s faild", jobid)
		fmt.Fprintf(w, err.Error())
		return
	}

	_, err = build.Stop()
	if err != nil {
		log.Errorf("stop build job:%s faild", jobid)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "true")
	return
}

func HandlerBuildConsoleOutput(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	jobid := q.Get(":jobid")
	number := q.Get(":number")
	fmt.Println(number)
	n, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Errorf("str to int64 faild")
		return
	}

	build, err := JenkinsClient.GetBuild(jobid, n)
	if err != nil {
		log.Errorf("get job:%s build faild", jobid)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output := build.GetConsoleOutput()
	fmt.Fprintf(w, output)
	return
}

func HandlerJobConfig(w http.ResponseWriter, r *http.Request) {
	jobid := r.URL.Query().Get(":jobid")
	job, err := JenkinsClient.GetJob(jobid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	config, err := job.GetConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("config:", config)
	w.Header().Set("Content-Type", "application/xml")
	fmt.Fprintf(w, config)
	return
}

/*
func HandlerCopyJob(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get(":from")
	newName := r.URL.Query().Get(":newname")

	job, err := JenkinsClient.GetJob(from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	j, err := job.Copy(newName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(j)
	if j != nil {
		fmt.Fprint(w, "true")
	} else {
		fmt.Fprint(w, "false")
	}
}
*/

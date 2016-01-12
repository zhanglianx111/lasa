package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bndr/gojenkins"
	"github.com/go-martini/martini"
	//"fmt"
)

func HandlerBuildConsoleOutput(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	/*
		number := params["number"]
		fmt.Println(number)
		_, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			log.Errorf("str to int64 faild")
			return
		}
	*/
	tBuildLog, err := template.ParseFiles("./views/buildlog.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	type buildlog struct {
		Log string
	}
	bLog := buildlog{}
	time.Sleep(1000 * time.Millisecond)
	jc := getJenkinsClient(r)
	latestId := getLatestBuildID(jc, jobid)
	if latestId == -1 {
		bLog.Log = "no log found, please Do Build Action"
		tBuildLog.Execute(w, bLog)
		return
	}
	build, err := jc.GetBuild(jobid, latestId)
	if err != nil {
		log.Errorf("get job:%s build faild", jobid)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output := build.GetConsoleOutput()
	bLog.Log = output
	tBuildLog.Execute(w, bLog)
	return
}

func getLatestBuildID(jc *gojenkins.Jenkins, jobname string) int64 {
	job, err := jc.GetJob(jobname)
	if err != nil {
		log.Errorf(err.Error())
		return -1
	}

	ids, err := job.GetAllBuildIds()
	if err != nil {
		log.Errorf(err.Error())
		return -1
	}
	/*
		for _, id := range ids {
			fmt.Println(id.Number)
		}
	*/
	if len(ids) != 0 {
		log.Debugf("%s latest id:%d", jobname, ids[0].Number)
		return ids[0].Number
	} else {
		return -1
	}
}

func HandlerGetBuildResult(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	/*
		number := params["number"]
		fmt.Println(number)
		n, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			log.Errorf("str to int64 faild")
			return
		}
	*/
	//time.Sleep(5000 * time.Millisecond)
	jc := getJenkinsClient(r)
	latestId := getLatestBuildID(jc, jobid)
	if latestId == -1 {
		return
	}
	build, err := jc.GetBuild(jobid, latestId)
	if err != nil {
		log.Errorf("get job:%s build faild", jobid)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rst := build.GetResult()
	fmt.Fprintf(w, rst)
	log.Debugf("build id: %d, build result:%s", latestId, rst)
	return
}

func HandlerStop(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	number := params["number"]
	fmt.Println(number)
	n, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Errorf("str to int64 faild")
		return
	}
	jc := getJenkinsClient(r)
	build, err := jc.GetBuild(jobid, n)
	if err != nil {
		log.Errorf("get job:%s build faild", jobid)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = build.Stop()
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Fprintf(w, "true")
	return
}

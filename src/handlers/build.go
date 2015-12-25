package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	//"net/http"
	//"fmt"
)

func HandlerBuildConsoleOutput(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	number := params["number"]
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

	fmt.Println(build.GetResult())
	output := build.GetConsoleOutput()
	fmt.Fprintf(w, output)
	return
}

func HandlerGetBuildResult(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	number := params["number"]
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

	rst := build.GetResult()
	fmt.Fprintf(w, rst)
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

	build, err := JenkinsClient.GetBuild(jobid, n)
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

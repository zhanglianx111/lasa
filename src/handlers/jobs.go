package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"utils"

	log "github.com/Sirupsen/logrus"
	"github.com/beevik/etree"
	simplejs "github.com/bitly/go-simplejson"
)

//job configuration path
/*
	./project/scm/userRemoteConfigs/hudson.plugins.git.UserRemoteConfig/url
	./project/description
	./project/scm/userRemoteConfigs/hudson.plugins.git.UserRemoteConfig/credentialsId
	./project/builders/com.cloudbees.dockerpublish.DockerBuilder/repoName
	./project/builders/com.cloudbees.dockerpublish.DockerBuilder/repoTag
	./project/builders/com.cloudbees.dockerpublish.DockerBuilder/skipPush
	./project/builders/hudson.tasks.Shell/command
*/
const (
	// 1
	Root = "./project/"
	// 2
	Scm      = "scm/"
	Builders = "builders/"
	// 3
	UsrRemoteConfigs                       = "userRemoteConfigs/"
	ComCloudbeesDockerpublishDockerBuilder = "com.cloudbees.dockerpublish.DockerBuilder/"
	// 4
	Branches                         = "branches/"
	HudsonPluginsGitBranchSpec       = "hudson.plugins.git.BranchSpec/"
	HudsonTasksShell                 = "hudson.tasks.Shell/"
	HudsonPluginsGitUserRemoteConfig = "hudson.plugins.git.UserRemoteConfig/"
	Registry                         = "registry/"
	Server                           = "server/"
	// 5
	Uri           = "uri"           //docker主机 ip:port
	CredentialsId = "credentialsId" //源代码仓库的用户名和密码
	Name          = "name"          //使用源代码的哪个分支，例如：master，dev等
	RepoName      = "repoName"      //image的repo信息
	RepoTag       = "repoTag"       //image的tag信息
	SkipPush      = "skipPush"      //是否跳过push操作，bool:ture or false
	SkipTagLatest = "skipTagLatest" //是否跳过使用latest最为tag，bool: true or false
	Command       = "command"       //使用execute shell的命令内容
	Url           = "url"           //docker registry address or 用户项目源代码仓库地址
	Description   = "description"   //用户项目的描述

	BaseCfg = "/Users/zhanglianxiang/workspace/jenkins_api/src/handlers/_tests/config.xml"
)

/*
type JobCfg struct {
	JobName     string `json:jobname`
	Description string `json:description`
	Scm         string `json:scm`
	Build       map[string]interface{}
}
*/
/*
用户自定义参数：
{
	"description": string, // default: ""
	"scm" : {
		"repositryurl": string, //用户项目的git地址
		"credentialsid": string, //用户usrname/passwd
		"branchestobuild": stirng, //用户想要build的分支名称
	}
    "builders": {
        "dockerbuildandpublish":{
            "repositryname": string, //用户自定义的在dockerregistry中的repo name
            "tag": string,             //images tag infomation
            "dockerhosturi": string, //default: http://dhub.yunpro.cn
			"dockerregitstryurl": string, //docker url
			"skippush": bool //是否跳过push到docker registry, defaule:false
        },
        "executeshell": {
			command: string // default: ""
        }
    }
}

*/

func HandlerCreateJob(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	jobid := params.Get(":jobid")
	js, err := simplejs.NewFromReader(r.Body)
	if err != nil {
		log.Errorf(err.Error())
		fmt.Fprintf(w, err.Error())
		return
	}
	defer r.Body.Close()

	cfg := parseCreateJobBody(js)
	if cfg == nil {
		log.Errorf("parse request body failde")
		fmt.Fprintf(w, "parse request body failde")
		return
	}
	fmt.Println(cfg)
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(BaseCfg); err != nil {
		log.Errorf("read job config.xml failed")
		fmt.Fprintf(w, "read job config.xml failed")
		return
	}
	// parse job config.xml
	updateJobConfigXml(doc, cfg)
	job_data, err := doc.WriteToString()
	if err != nil {
		log.Errorf("write to string failed")
		fmt.Fprintf(w, "write to string failed")
		return
	}

	job, err := JenkinsClient.CreateJob(job_data, jobid)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	//creating job is ok
	log.Debugf("job base:%s, job raw:", job.Base, job.Raw)
	data := map[string]string{"name": job.Raw.Name} // "description": job.Raw.Description,
	// "displayName": job.Raw.DisplayName,
	// "url":         job.Raw.URL

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

func parseCreateJobBody(js *simplejs.Json) map[string]string {
	m := make(map[string]string)

	if desc, err := js.Get("description").String(); err != nil {
		return nil
	} else {
		m["desc"] = desc
	}

	// scm
	scm := js.Get("scm")
	if scm == nil {
		return nil
	}

	if repositryrUrl, err := scm.Get("repositryurl").String(); err != nil {
		return nil
	} else {
		m["repositryurl"] = repositryrUrl
	}

	if credentialsId, err := scm.Get("credentialsid").String(); err != nil {
		return nil
	} else {
		m["credentialsid"] = credentialsId
	}

	if branchesToBuild, err := scm.Get("branchestobuild").String(); err != nil {
		return nil
	} else {
		m["branchestobuild"] = branchesToBuild
	}
	// builders
	bdrs := js.Get("builders").Get("dockerbuildandpublish")
	if bdrs == nil {
		return nil
	}

	if repoName, err := bdrs.Get("repositryname").String(); err != nil {
		return nil
	} else {
		m["repositryname"] = repoName
	}

	if tag, err := bdrs.Get("tag").String(); err != nil {
		return nil
	} else {
		m["tag"] = tag
	}

	if dockerRegistryUrl, err := bdrs.Get("dockerregitstryurl").String(); err != nil {
		return nil
	} else {
		m["dockerregitstryurl"] = dockerRegistryUrl
	}

	if dockerHostUri, err := bdrs.Get("dockerhosturi").String(); err != nil {
		return nil
	} else {
		m["dockerhosturi"] = dockerHostUri
	}

	if skipPush, err := bdrs.Get("skippush").String(); err != nil {
		return nil
	} else {
		m["skippush"] = skipPush
	}

	if cmd, err := js.Get("builders").Get("executeshell").Get("command").String(); err != nil {
		return nil
	} else {
		m["command"] = cmd
	}

	return m
}

func updateJobConfigXml(doc *etree.Document, cfg map[string]string) {
	eDesc := doc.FindElement(Root + Description)
	eDesc.SetText(cfg["desc"])

	eRepoURL := doc.FindElement(Root + Scm + UsrRemoteConfigs + HudsonPluginsGitUserRemoteConfig + Url)
	eRepoURL.SetText(cfg["repositryurl"])

	eCredentialsid := doc.FindElement(Root + Scm + UsrRemoteConfigs + HudsonPluginsGitUserRemoteConfig + CredentialsId)
	eCredentialsid.SetText(cfg["credentialsid"])

	eBranchesToBuild := doc.FindElement(Root + Scm + Branches + HudsonPluginsGitBranchSpec + Name)
	eBranchesToBuild.SetText(cfg["branchestobuild"])

	eRepoName := doc.FindElement(Root + Builders + ComCloudbeesDockerpublishDockerBuilder + RepoName)
	eRepoName.SetText(cfg["repositryname"])

	eTag := doc.FindElement(Root + Builders + ComCloudbeesDockerpublishDockerBuilder + RepoTag)
	eTag.SetText(cfg["tag"])

	eDockerHostUri := doc.FindElement(Root + Builders + ComCloudbeesDockerpublishDockerBuilder + Server + Uri)
	eDockerHostUri.SetText(cfg["dockerhosturi"])

	eDockerRegistryUrl := doc.FindElement(Root + Builders + ComCloudbeesDockerpublishDockerBuilder + Registry + Url)
	eDockerRegistryUrl.SetText(cfg["dockerregitstryurl"])

	eSkipPush := doc.FindElement(Root + Builders + ComCloudbeesDockerpublishDockerBuilder + SkipPush)
	eSkipPush.SetText(cfg["skippush"])

	eCmd := doc.FindElement(Root + Builders + HudsonTasksShell + Command)
	eCmd.SetText(cfg["command"])

}

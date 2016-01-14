package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"utils"

	log "github.com/Sirupsen/logrus"
	"github.com/beevik/etree"
	simplejs "github.com/bitly/go-simplejson"
	"github.com/clbanning/mxj"
	"github.com/go-martini/martini"
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


func HandlerCreateJob(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	jobid := params.Get(":jobid")
	fmt.Println("job id is: ", jobid)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	js, err := simplejs.NewFromReader(r.Body)
	if err != nil {
		log.Errorf(err.Error())
		log.Errorf("read request body failed")
		fmt.Fprintf(w, err.Error())
		return
	}
	defer r.Body.Close()

	cfg := parseCreateJobBody(js)
	if cfg == nil {
		log.Errorf("parse request body failed")
		fmt.Fprintf(w, "parse request body failed")
		return
	}
	/*
		doc := etree.NewDocument()
		if err := doc.ReadFromFile(BaseCfg); err != nil {
			log.Errorf("read job config.xml failed")
			fmt.Fprintf(w, "read job config.xml failed")
			return
		}
*/
/*
	doc := JobConfig.Copy()
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
		log.Errorf("create job: %s failed", jobid)
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

	w.Write(jsonData)
}
*/
func HandlerCreateJob(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tCreateJob, err := template.ParseFiles("views/createJob.html")
		if err != nil {
			fmt.Println(w, err.Error())
			return
		}
		tCreateJob.Execute(w, nil)
	} else {
		js, err := simplejs.NewFromReader(r.Body)
		if err != nil {
			log.Errorf(err.Error())
			log.Errorf("read request body failed")
			fmt.Fprintf(w, err.Error())
			return
		}
		defer r.Body.Close()
		fmt.Println(js)
		jobid, cfg := parseCreateJobBody(js)
		if cfg == nil {
			log.Errorf("parse request body failed")
			fmt.Fprintf(w, "parse request body failed")
			return
		}

		doc := JobConfig.Copy()
		// parse job config.xml
		updateJobConfigXml(doc, cfg)
		job_data, err := doc.WriteToString()
		if err != nil {
			log.Errorf("write to string failed")
			fmt.Fprintf(w, "write to string failed")
			return
		}
		cookie, _ := r.Cookie("sessionId")
		jc := getJenkinsClient(cookie.Value)
		job, err := jc.CreateJob(job_data, jobid)
		if err != nil {
			log.Errorf("create job: %s failed", jobid)
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

		w.Write(jsonData)
	}
}
func HandlerDeleteJob(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	log.Debugf("delete job:%s", jobid)
	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	_, err := jc.DeleteJob(jobid)
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

func HandlerDisableJob(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	log.Debugf("disable job: %s", jobid)
	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	job, err := jc.GetJob(jobid)
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

func HandlerEnableJob(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	log.Debugf("enable job:%s", jobid)
	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	job, err := jc.GetJob(jobid)
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

func HandlerRenameJob(params martini.Params, w http.ResponseWriter, r *http.Request) {
	oldName := params["oldjobid"]
	newName := params["newjobid"]
	if oldName == "" || newName == "" {
		fmt.Fprintf(w, "oldname or newname is empty")
		return
	}

	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	job := jc.RenameJob(oldName, newName)
	if job == nil {
		fmt.Fprintf(w, "job rename faild")
	}

	fmt.Fprintf(w, job.Base)
	return
}

func HandlerGetJob(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params[":jobid"]

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	job, err := jc.GetJob(jobid)
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

	w.Write(jsonData)
}

func HandlerGetAllJobs(w http.ResponseWriter, r *http.Request) {
	log.Debugf("%s %s", r.Method, r.URL)
	var jobsData []map[string]interface{}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cookie, _ := r.Cookie("sessionId")
	//sess, _ := GlobalSessions.SessionStart(w, r)
	sess, _ := GlobalSessions.GetSessionStore(cookie.Value)
	log.Debugf("user:%s", sess.Get("user"))
	fmt.Println(sess)
	log.Debugf("sid: %s", sess.SessionID())
	jc := getJenkinsClient(cookie.Value)
	fmt.Println("jc:", jc)
	jobs, _ := jc.GetAllJobs()
	for _, job := range jobs {
		jobData := utils.AnalysisJob(job.Raw)
		jobsData = append(jobsData, jobData)
	}

	jsonData, err := json.Marshal(jobsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

func HandlerGetAllBuildIds(params martini.Params, w http.ResponseWriter, r *http.Request) {
	bs := []map[string]interface{}{}
	jobid := params["jobid"]
	if jobid == "" {
		fmt.Fprintf(w, "jobid is empty")
		return
	}

	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	job, err1 := jc.GetJob(jobid)
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

func HandlerJobRunning(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	if jobid == "" {
		fmt.Fprintf(w, "jobid is empty")
		return
	}

	cookie, _ := r.Cookie("sessionid")
	jc := getJenkinsClient(cookie.Value)
	job, err := jc.GetJob(jobid)
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

func HandlerBuildJob(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	job, err := jc.GetJob(jobid)
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

	b, err := jc.BuildJob(jobid)
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

func HandlerStopBuild(params martini.Params, w http.ResponseWriter, r *http.Request) {
	jobid := params["jobid"]
	number := params["number"]
	n, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Errorf("string to int64 faild")
		return
	}
	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	build, err := jc.GetBuild(jobid, n)
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

func HandlerJobConfig(params martini.Params, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		jobid := params["jobid"]
		cookie, _ := r.Cookie("sessionId")
		jcc := getJenkinsClient(cookie.Value)
		job, err := jcc.GetJob(jobid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// config is a xml format
		config, err := job.GetConfig()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		m, err := mxj.NewMapXml([]byte(config))
		if err != nil {
			log.Errorf(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		kpv := map[string]*mxj.LeafNode{
			"description":  &mxj.LeafNode{},
			"uri":          &mxj.LeafNode{}, // docker host
			"url4git":      &mxj.LeafNode{}, // project git address
			"url4registry": &mxj.LeafNode{}, // registry address
			//"usr":         &mxj.LeafNode{},
			//"passwd":      &mxj.LeafNode{},
			"name":     &mxj.LeafNode{}, // branches name
			"repoName": &mxj.LeafNode{},
			"repoTag":  &mxj.LeafNode{},
			"skipPush": &mxj.LeafNode{},
			"command":  &mxj.LeafNode{},
		}
		paths := m.PathsForKey("url")
		for _, p := range paths {
			if strings.Contains(p, "UserRemoteConfig") {
				kpv["url4git"].Path = p
			} else if strings.Contains(p, "registry") {
				kpv["url4registry"].Path = p
			} else {
				log.Warnf("not founc url path")
			}
		}

		for key, _ := range kpv {
			if key != "url4git" && key != "url4registry" {
				kpv[key].Path = m.PathForKeyShortest(key)
			}
		}

		leafNodes := m.LeafNodes()
		for _, v := range leafNodes {
			for k, vv := range kpv {
				if vv.Path == v.Path {
					kpv[k].Value = v.Value
				}
			}
		}

		type jobConfig struct {
			Name        string // job name
			Desc        string // job description
			DockerHost  string // docker host
			Username    string
			Passwd      string
			Branches    string // branches name
			RepoName    string // repositry name
			RepoTag     string // repositry tag
			Skippush    string // skip push
			RegistryUrl string // registry url
			GitUrl      string
			Command     string // command
		}
		var jc jobConfig
		jc.Name = jobid
		jc.Desc = reflect.ValueOf(kpv["description"].Value).String()
		jc.DockerHost = reflect.ValueOf(kpv["uri"].Value).String()
		jc.Username = "username"
		jc.Passwd = "password"
		jc.Branches = reflect.ValueOf(kpv["name"].Value).String()
		jc.RepoName = reflect.ValueOf(kpv["repoName"].Value).String()
		jc.RepoTag = reflect.ValueOf(kpv["repoTag"].Value).String()

		if reflect.ValueOf(kpv["skipPush"].Value).String() == "true" {
			jc.Skippush = "checked"
		}
		jc.RegistryUrl = reflect.ValueOf(kpv["url4registry"].Value).String()
		jc.GitUrl = reflect.ValueOf(kpv["url4git"].Value).String()
		jc.Command = reflect.ValueOf(kpv["command"].Value).String()

		tJobConfig, err := template.ParseFiles("./views/jobConfig1.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		tJobConfig.Execute(w, jc)
		return
	} else {
		fmt.Fprintf(w, "developing...")
	}
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

func parseCreateJobBody(js *simplejs.Json) (string, map[string]string) {
	m := make(map[string]string)
	jobname, err := js.Get("jobname").String()
	if err != nil {
		log.Errorf("parse \"jobname\" failed")
		return "", nil
	}
	/*
		else {
			m["jobname"] = jobname
			fmt.Println(jobname)
		}*/
	if desc, err := js.Get("description").String(); err != nil {
		log.Errorf("parse \"description\" failed")
		return "", nil
	} else {
		m["desc"] = desc
	}

	if repositryrUrl, err := js.Get("repositryurl").String(); err != nil {
		//log.Errorf("parse \"repositryurl\" failed")
		log.Errorf(err.Error())
		return "", nil
	} else {
		m["repositryurl"] = repositryrUrl
	}
	/*
		if credentialsId, err := js.Get("credentialsid").String(); err != nil {
			log.Errorf("parse \"credentialsid\" failed")
			return nil
		} else {
			m["credentialsid"] = credentialsId
		}
	*/
	if usr, err := js.Get("username").String(); err != nil {
		log.Errorf("parse \"username\" failed")
		return "", nil
	} else {
		m["username"] = usr
	}

	if passwd, err := js.Get("password").String(); err != nil {
		log.Errorf("parse \"password\" failed")
		return "", nil
	} else {
		m["password"] = passwd
	}

	if branchesToBuild, err := js.Get("branches").String(); err != nil {
		log.Errorf("parse \"branches\" failed")
		return "", nil
	} else {
		m["branchestobuild"] = branchesToBuild
	}
	/*
		// builders
		bdrs := js.Get("builders").Get("dockerbuildandpublish")
		if bdrs == nil {
			log.Errorf("parse \"dockerbuildandpublish\" failed")
			return nil
		}
	*/
	if repoName, err := js.Get("repositryname").String(); err != nil {
		log.Errorf("parse \"repositryname\" failed")
		return "", nil
	} else {
		m["repositryname"] = repoName
	}

	if tag, err := js.Get("tag").String(); err != nil {
		log.Errorf("parse \"tag\" failed")
		return "", nil
	} else {
		m["tag"] = tag
	}

	if dockerRegistryUrl, err := js.Get("dockerregistryurl").String(); err != nil {
		log.Errorf("parse \"dockerregistryurl\" failed")
		return "", nil
	} else {
		m["dockerregistryurl"] = dockerRegistryUrl
	}

	if dockerHostUri, err := js.Get("dockerhosturi").String(); err != nil {
		log.Errorf("parse \"dockerhosturi\" failed")
		return "", nil
	} else {
		m["dockerhosturi"] = dockerHostUri
	}

	if skipPush, err := js.Get("skippush").String(); err != nil {
		log.Errorf("parse \"skippush\" failed")
		return "", nil
	} else {
		m["skippush"] = skipPush
	}

	if cmd, err := js.Get("command").String(); err != nil {
		log.Errorf("parse \"command\" failed")
		return "", nil
	} else {
		m["command"] = cmd
	}

	return jobname, m
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
	eDockerRegistryUrl.SetText(cfg["dockerregistryurl"])

	eSkipPush := doc.FindElement(Root + Builders + ComCloudbeesDockerpublishDockerBuilder + SkipPush)
	eSkipPush.SetText(cfg["skippush"])

	eCmd := doc.FindElement(Root + Builders + HudsonTasksShell + Command)
	eCmd.SetText(cfg["command"])

}

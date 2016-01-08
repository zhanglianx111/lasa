package handlers

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/session"
	"github.com/bndr/gojenkins"
	//"github.com/zhanglianx111/gojenkins"
	"github.com/beevik/etree"
)

var JenkinsClient map[string]*gojenkins.Jenkins
var GlobalSessions *session.Manager
var JobConfig *etree.Document
var BaseCfg = "/Users/zhanglianxiang/workspace/jenkins_api/src/handlers/_tests/config.xml"

type JenkinsInfo struct {
	Jobs      []string `json:jobs`
	Mode      string   `json:mode`
	NodesName []string `json:nodeName`
}

func init() {
	log.SetLevel(log.DebugLevel)
	// do a deep copy for etree of job config.xml
	JobConfig = etree.NewDocument()
	if err := JobConfig.ReadFromFile(BaseCfg); err != nil {
		log.Errorf(err.Error())
		return
	}

	// session uses memory
	GlobalSessions, _ = session.NewManager(
		"cookie", `{"cookieName":"sessionId","enableSetCookie":true,"gclifetime":30,"ProviderConfig":"{\"cookieName\":\"sessionId\",\"securityKey\":\"beegocookiehashkey\"}"}`)
	go GlobalSessions.GC()
	JenkinsClient = make(map[string]*gojenkins.Jenkins)
	/*
		for {
			JenkinsClient = getJenkinsClient()
			if JenkinsClient == nil {
				time.Sleep(10)
				continue
			} else {
				break
			}
		}
	*/
}

/*
func getJenkinsClient() *gojenkins.Jenkins {
	var jenkinsHost, jenkinsPort string
	if os.Getenv("ENVIRONMENT") == "production" {
		jenkinsHost = os.Getenv("JENKINS_HOST")
		jenkinsPort = os.Getenv("JENKINS_PORT")
		if jenkinsHost == "" || jenkinsPort == "" {
			log.Errorf("jenkinsHost:%s, jenkinsPort:%s", jenkinsHost, jenkinsPort)
			return nil
		}
	} else {
		jenkinsHost = "10.10.11.207"
		jenkinsPort = "8080"
	}

	url := "http://" + jenkinsHost + ":" + jenkinsPort
	jenkins, err := gojenkins.CreateJenkins(url, "admin", "admin").Init()
	if err != nil {
		log.Errorf("connecting jenkins server:%s:%s failed with error:%s", jenkinsHost, jenkinsPort, err)
		return nil
	}
	log.Infof("connect jenkins server:%s:%s is OK!", jenkinsHost, jenkinsPort)
	return jenkins
}
*/
func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Error method:", r.Method)
	}
	http.Redirect(w, r, "/login", http.StatusFound)
	//	fmt.Fprintf(w, "jenkins rest api server")
}

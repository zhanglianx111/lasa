package handlers

import (
	"fmt"

	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bndr/gojenkins"
	//"github.com/zhanglianx111/gojenkins"
	"github.com/beevik/etree"
)

var JenkinsClient *gojenkins.Jenkins
var JobConfig *etree.Document
var BaseCfg = "/Users/zhanglianxiang/workspace/jenkins_api/src/handlers/_tests/config.xml"

type JenkinsInfo struct {
	Jobs      []string `json:jobs`
	Mode      string   `json:mode`
	NodesName []string `json:nodeName`
}

func Init() {
	log.SetLevel(log.DebugLevel)
	// do a deep copy for etree of job  config.xml
	JobConfig = etree.NewDocument()
	if err := JobConfig.ReadFromFile(BaseCfg); err != nil {
		return
	}
	for {
		JenkinsClient = GetJenkinsClient()
		if JenkinsClient == nil {
			time.Sleep(10)
			continue
		} else {
			break
		}
	}
}

func GetJenkinsClient() *gojenkins.Jenkins {
	var jenkinsHost, jenkinsPort string
	if os.Getenv("ENVIRONMENT") == "production" {
		jenkinsHost = os.Getenv("JENKINS_HOST")
		jenkinsPort = os.Getenv("JENKINS_PORT")
		if jenkinsHost == "" || jenkinsPort == "" {
			return nil
		}
	} else {
		jenkinsHost = "10.10.11.111"
		jenkinsPort = "8080"
	}

	url := "http://" + jenkinsHost + ":" + jenkinsPort
	jenkins, err := gojenkins.CreateJenkins(url, "admin", "admin").Init()
	if err != nil {
		log.Errorf("connecting jenkins server:%s:%s failed br %s", jenkinsHost, jenkinsPort, err)
		return nil
	}
	log.Infof("connecting jenkins server:%s:%s is OK!", jenkinsHost, jenkinsPort)
	return jenkins
}

func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is jenkins api server!!!")
}

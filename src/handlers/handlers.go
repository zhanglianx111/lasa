package handlers

import (
	"fmt"

	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bndr/gojenkins"
	//"github.com/zhanglianx111/gojenkins"
)

var JenkinsClient *gojenkins.Jenkins

type JenkinsInfo struct {
	Jobs      []string `json:jobs`
	Mode      string   `json:mode`
	NodesName []string `json:nodeName`
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

func Init() {
	log.SetLevel(log.DebugLevel)
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

func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is jenkins api server!!!")
}

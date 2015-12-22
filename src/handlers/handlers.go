package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bndr/gojenkins"
	//"github.com/zhanglianx111/gojenkins"
	"github.com/beevik/etree"
	"github.com/tsuru/config"
)

var JenkinsClient *gojenkins.Jenkins
var JobConfig *etree.Document
var BaseCfg = "/Users/zhanglianxiang/workspace/jenkins_api/src/handlers/_tests/config.xml"
var ConfigurationPath = "config/config.yml"

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
		return
	}
	for {
		JenkinsClient = getJenkinsClient()
		if JenkinsClient == nil {
			time.Sleep(10)
			continue
		} else {
			break
		}
	}
}

func getJenkinsClient() *gojenkins.Jenkins {
	var jenkinsHost string
	var jenkinsPort int
	configAbsPath, err := filepath.Abs(ConfigurationPath)
	if err != nil {
		log.Errorf(err.Error())
		return nil
	}
	err = config.ReadAndWatchConfigFile(configAbsPath)
	if err != nil {
		log.Errorf("read app configuration file failed", err.Error())
		return nil
	}
	jenkinsHost, err = config.GetString("server:host")
	if err != nil {
		log.Errorf(err.Error())
		return nil
	}

	jenkinsPort, err = config.GetInt("server:port")
	if err != nil {
		log.Errorf(err.Error())
		return nil
	}

	url := "http://" + jenkinsHost + ":" + strconv.Itoa(jenkinsPort)
	jenkins, err := gojenkins.CreateJenkins(url, "admin", "admin").Init()
	if err != nil {
		log.Errorf("connecting jenkins server:%s:%d failed with error:%s", jenkinsHost, jenkinsPort, err)
		return nil
	}
	log.Infof("connect jenkins server:%s:%d is OK!", jenkinsHost, jenkinsPort)
	return jenkins
}

func HandlerDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Error method:", r.Method)
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"utils"
)

type Node struct {
	Name string `json:"name"`
}

/*
	return [map, map]
*/

func HandlerGetAllNodes(w http.ResponseWriter, r *http.Request) {
	var nodes []Node

	allNodes, err := JenkinsClient.GetAllNodes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, node := range allNodes {
		nodes = append(nodes, Node{node.Base})
	}

	jsonData, err := json.Marshal(nodes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

/*
	return map{"name": "nodename"}
*/
func HandlerGetNode(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	nodeName := params.Get(":nodeid")

	node, err := JenkinsClient.GetNode(nodeName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("node Raw:", node.Raw)

	nodeData := utils.AnalysisNode(node.Raw)
	jsonData, err := json.Marshal(nodeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandlerAddNode(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	nodeName := params.Get(":nodeid")
	description := params.Get(":description")
	remotefs := params.Get(":remotefs")
	num := params.Get(":numexecutors")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error")
		return
	}

	var launcher map[string]string
	err = json.Unmarshal(body, &launcher)
	if err != nil {
		fmt.Fprint(w, "convert json format failed")
		return
	}

	sshLauncher := map[string]string{
		"method":               launcher["method"],
		"host":                 launcher["host"],
		"port":                 launcher["port"],
		"credentialsId":        launcher["credentialsId"],
		"jvmOptions":           launcher["jvmOptions"],
		"javaPath":             launcher["javaPath"],
		"prefixStartSlaveCmd":  launcher["efixStartSlaveCmd"],
		"suffixStartSlaveCmd":  launcher["suffixStartSlaveCmd"],
		"maxNumRetries":        launcher["maxNumRetries"],
		"retryWaitTime":        launcher["retryWaitTime"],
		"lanuchTimeoutSeconds": launcher["lanuchTimeoutSeconds"],
	}

	fmt.Println(sshLauncher)
	numexecutors, err := strconv.Atoi(num)
	if err != nil {
		fmt.Fprintf(w, "Atoi faild")
		return
	}

	rf := strings.Split(remotefs, " ")
	rfs := "/" + rf[0] + "/" + rf[1] + "/" + rf[2]
	fmt.Println(params.Get(":method"))
	node, err := JenkinsClient.CreateNode(nodeName, numexecutors, description, rfs, sshLauncher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(node)
	fmt.Fprintf(w, "true")

}

func HandlerDeleteNode(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	nodeName := params.Get(":nodeid")

	node, err := JenkinsClient.GetNode(nodeName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = node.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "true")
}

func HandlerNodeToggleTemporarilyOffline(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")
	if nodeid == "" {
		fmt.Fprint(w, "nodeid is empty")
		return
	}

	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err1 := node.ToggleTemporarilyOffline()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	if b {
		fmt.Fprint(w, "true")
	} else {
		fmt.Fprint(w, "false")
	}
}

func HandlerSetOnline(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")
	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err1 := node.SetOnline()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	if b {
		fmt.Fprint(w, "true")
	} else {
		fmt.Fprint(w, "false")
	}
}

func HandlerSetOffline(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")
	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err1 := node.SetOffline()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	if b {
		fmt.Fprint(w, "true")
	} else {
		fmt.Fprint(w, "false")
	}
}

func HandlerNodeInfo(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")
	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	info, err1 := node.Info()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(info)
	return
}

func HandlerIsJnlpAgent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlerIsJnlpAgent")
}

func HandlerIsIdle(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")
	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err1 := node.IsIdle()
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	if b {
		fmt.Fprint(w, "ture")
	} else {
		fmt.Fprint(w, "false")
	}
}

func HandlersLaunchNode(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")
	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = node.LaunchNodeBySSH()
	if err != nil {
		fmt.Fprintf(w, "lanuch node:%s failed", nodeid)
		return
	}
	fmt.Fprintf(w, "true")
	return
}

func HandlerDisconnect(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")

	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = node.Disconnect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "true")
	return
}

func HandlerGetLogText(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")
	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("+++++")
		return
	}

	log, err := node.GetLogText()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, log)
	return
}

func HandlerIsOnline(w http.ResponseWriter, r *http.Request) {
	nodeid := r.URL.Query().Get(":nodeid")
	node, err := JenkinsClient.GetNode(nodeid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = node.IsOnline()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "log")
	return
}

package handlers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/bndr/gojenkins"
	"github.com/go-martini/martini"
	//"github.com/zhanglianx111/gojenkins"
	"net/http"
)

func HandlerCreateView(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	viewName := params.Get(":name")
	viewType := params.Get(":type")
	fmt.Println(viewName, viewType)
	jc := getJenkinsClient(r)
	view, err := jc.CreateView(viewName, gojenkins.LIST_VIEW)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	} else {
		fmt.Fprintf(w, "true")
		fmt.Println(view)
	}
	return
}

func HandlerDeleteView(w http.ResponseWriter, r *http.Request) {
	// TODO

}

func HandlerViewAddJob(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	viewName := params.Get(":name")
	jobName := params.Get(":jobid")
	fmt.Println(viewName, jobName)
	if viewName == "" || jobName == "" {
		fmt.Fprintf(w, "viewName or jobid is empty")
		return
	}

	jc := getJenkinsClient(r)
	view, err := jc.GetView(viewName)
	fmt.Println(view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := view.AddJob(jobName)
	if !b {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	/*
		jsonData, err := json.Marshal(map[string]string{"jobName": jobName})
		if err != nil {
			fmt.Fprintf(w, "json marshal faild")
			return
		}
	*/
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "true")
}

func HandlerViewDeleteJob(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	viewName := params.Get(":name")
	jobName := params.Get(":jobid")
	fmt.Println(viewName, jobName)
	if viewName == "" || jobName == "" {
		fmt.Fprintf(w, "viewName or jobid is empty")
		return
	}

	jc := getJenkinsClient(r)
	view, err := jc.GetView(viewName)
	fmt.Println(view)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, _ := view.DeleteJob(jobName)
	if !b {
		fmt.Fprintf(w, "false")
	} else {
		fmt.Fprintf(w, "true")
	}
}

/*
	return []map[string]interface{}
*/
func HandlerGetAllViews(w http.ResponseWriter, r *http.Request) {
	var allViews []map[string]interface{}

	jc := getJenkinsClient(r)
	views, err := jc.GetAllViews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(views); i++ {
		allViews = append(allViews, map[string]interface{}{"name": views[i].GetName(), "url": views[i].GetUrl(), "jobs": views[i].GetJobs()})
		fmt.Println(allViews)
	}

	jsonData, err := json.Marshal(allViews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandlerGetView(params martini.Params, w http.ResponseWriter, r *http.Request) {
	log.Debugf("%s %s", r.Method, r.URL)
	cookie, _ := r.Cookie("sessionId")
	sess, _ := GlobalSessions.GetSessionStore(cookie.Value)
	log.Debugf("user:%s", sess.Get("user"))
	fmt.Println(sess)
	log.Debugf("sid: %s", sess.SessionID())
	viewName := params["viewid"]
	user := reflect.ValueOf(sess.Get("user")).String()
	if strings.Compare(user, "admin") == 0 {
		viewName = "All"
	}
	if viewName == "" {
		fmt.Fprintf(w, "params(view) is empty")
		return
	}

	jc := getJenkinsClient(r)
	view, err := jc.GetView(viewName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(map[string]interface{}{"name": view.GetName(), "url": view.GetUrl(), "description": view.GetDescription(), "jobs": view.GetJobs()})
	if err != nil {
		fmt.Fprintf(w, "Marshal faild")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

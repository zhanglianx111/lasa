package handlers

import (
	//"github.com/bndr/gojenkins"
	"encoding/json"
	"fmt"
	"github.com/zhanglianx111/gojenkins"
	"net/http"
)

func HandlerCreateView(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	viewName := params.Get(":name")
	viewType := params.Get(":type")
	fmt.Println(viewName, viewType)
	view, err := JenkinsClient.CreateView(viewName, gojenkins.LIST_VIEW)
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

	view, err := JenkinsClient.GetView(viewName)
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

	view, err := JenkinsClient.GetView(viewName)
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

	views, err := JenkinsClient.GetAllViews()
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

func HandlerGetView(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	viewName := params.Get(":viewid")
	if viewName == "" {
		fmt.Fprintf(w, "params(view) is empty")
		return
	}

	view, err := JenkinsClient.GetView(viewName)
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

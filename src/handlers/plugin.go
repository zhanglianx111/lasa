package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
type plugin struct {

}
*/
func HandlerGetPlugins(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	d := params.Get(":depth")
	depth, _ := strconv.Atoi(d)
	fmt.Println(depth)
	cookie, _ := r.Cookie("sessionId")
	jc := getJenkinsClient(cookie.Value)
	plugins, err := jc.GetPlugins(depth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	plgs := []map[string]interface{}{}
	p := map[string]interface{}{}
	for _, plg := range plugins.Raw.Plugins {
		//plgs = append(plgs, p.LongName)
		p = map[string]interface{}{"name": plg.LongName,
			"active":  plg.Active,
			"version": plg.Version,
			"enabled": plg.Enabled,
			"url":     plg.URL}
		plgs = append(plgs, p)
		//fmt.Println(p.LongName)
	}
	jsonData, err := json.Marshal(plgs)
	if err != nil {
		fmt.Fprintf(w, "json marshal failed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

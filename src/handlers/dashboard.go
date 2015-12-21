package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func HandlerDashboard(w http.ResponseWriter, r *http.Request) {
	if t, err := template.ParseFiles("./views/index.html"); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	} else {
		t.Execute(w, nil)
	}
}

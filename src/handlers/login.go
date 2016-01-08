package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/login.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	} else {
		fmt.Println(r.PostFormValue("user"))
		user := r.PostFormValue("user")
		passwd := r.PostFormValue("passwd")
		fmt.Println(user)
		fmt.Println(passwd)
		if user == "zlx" && passwd == "zlx" {
			fmt.Fprintf(w, "success")
		} else {
			w.WriteHeader(401)
			fmt.Fprintf(w, "please checkout your email and password!")
		}
		/*
			if len(email) == 0 || len(passwd) == 0 {
				fmt.Fprintf(w, "login failure, please checkout your email and password!")
				return
			}
			// TODO authrezition
			// show client dashboard page
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			tDashboard, err := template.ParseFiles("./views/dashboard.html")
			if err != nil {
				fmt.Fprintf(w, err.Error())
				return
			}
			tDashboard.Execute(w, nil)
		*/
	}
}

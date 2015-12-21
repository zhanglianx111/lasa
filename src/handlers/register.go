package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HandlerRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./views/register.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	} else {
		fmt.Println(r.PostFormValue("email"))
		email := r.PostFormValue("email")
		passwd := r.PostFormValue("passwd")
		if email == "zlx@zlx.com" && passwd == "zlx" {
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

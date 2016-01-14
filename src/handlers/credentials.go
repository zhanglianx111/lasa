package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
)

func HandlerCredentials(params martini.Params, w http.ResponseWriter, r *http.Request) {
	user := params["user"]
	passwd := params["passwd"]
	scope := "GLOBAL"
	description := ""
	id := ""
	/*
		cookie, _ := r.Cookie("sessionId")
		sess, _ := GlobalSessions.GetSessionStore(cookie.Value)
		fmt.Println(sess.Get("user"))
	*/
	//jc := getJenkinsClient(cookie.Value)
	jc := getJenkinsClient("dd57e1c3520cbda92e6faf93eb261adc")
	err := jc.CreateCredential(scope, user, passwd, description, id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "success")

}

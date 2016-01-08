package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/bndr/gojenkins"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./views/login.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		w.Header().Set("Content-Type", "text/html")
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
	} else {
		sess, err := GlobalSessions.SessionStart(w, r)
		if err != nil {
			log.Warnf("session err: %s", err.Error())
		}
		defer sess.SessionRelease(w)

		sess.Set("user", r.PostFormValue("user"))
		user := r.PostFormValue("user")
		passwd := r.PostFormValue("passwd")
		fmt.Println("--seesionid:", sess.SessionID())
		fmt.Println("--username: ", sess.Get("user"))
		fmt.Println("passwd: ", passwd)
		//sess.SessionRelease(w)
		fmt.Println("==seesionid:", sess.SessionID())
		//if user == "zlx" && passwd == "zlx" {
		jc := addJenkinsClient(user, passwd)
		if jc == nil {
			log.Errorf("%s get jenkins cleint failed", user)
			fmt.Fprintf(w, "failuer")
			return
		}
		JenkinsClient[sess.SessionID()] = jc
		fmt.Println(JenkinsClient)
		fmt.Fprintf(w, "success")
		/*
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

func addJenkinsClient(user, passwd string) *gojenkins.Jenkins {
	var jenkinsHost, jenkinsPort string
	if os.Getenv("ENVIRONMENT") == "production" {
		jenkinsHost = os.Getenv("JENKINS_HOST")
		jenkinsPort = os.Getenv("JENKINS_PORT")
		if jenkinsHost == "" || jenkinsPort == "" {
			log.Errorf("jenkinsHost:%s, jenkinsPort:%s", jenkinsHost, jenkinsPort)
			return nil
		}
	} else {
		jenkinsHost = "10.10.11.207"
		jenkinsPort = "8080"
	}

	url := "http://" + jenkinsHost + ":" + jenkinsPort
	jenkins, err := gojenkins.CreateJenkins(url, user, passwd).Init()
	if err != nil {
		log.Errorf("user:%s connecting jenkins server:%s:%s failed with error:%s", user, jenkinsHost, jenkinsPort, err)
		return nil
	}
	log.Infof("user:%s connect jenkins server:%s:%s is OK!", user, jenkinsHost, jenkinsPort)
	return jenkins
}

func getJenkinsClient(r *http.Request) *gojenkins.Jenkins {
	sid, _ := r.Cookie("sessionId")
	log.Debugf("session id:%s", sid.Value)
	s, _ := GlobalSessions.GetSessionStore(sid.Value)
	fmt.Println(s.SessionID())
	jc := JenkinsClient[s.SessionID()]
	fmt.Println(jc)
	return jc
}

func HandlerLogout(w http.ResponseWriter, r *http.Request) {
	log.Debugf("%s %s", r.Method, r.URL)
	GlobalSessions.SessionDestroy(w, r)
	http.Redirect(w, r, "/login", http.StatusFound)
}

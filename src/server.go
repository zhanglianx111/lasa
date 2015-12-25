package main

import (
	"net/http"
	"router"

	"github.com/go-martini/martini"
)

func main() {
	//	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.New()
	// for static files service to dir static
	m.Use(martini.Static("static"))
	// for router
	m.MapTo(router.Routers, (*martini.Routes)(nil))
	m.Action(router.Routers.Handle)
	http.Handle("/", m)
	// run app at port 3000
	m.RunOnAddr("10.10.16.42:3000")
}

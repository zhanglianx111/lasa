package main

import (
	"net/http"
	"router"

	"github.com/go-martini/martini"
)

func main() {
	//	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.New()
	m.MapTo(router.Routers, (*martini.Routes)(nil))
	m.Action(router.Routers.Handle)
	http.Handle("/", m)
	m.RunOnAddr(":3000")
}

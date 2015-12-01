package main

import (
	"router"
	"net/http"
	"runtime"
)


func Init() {
	router.Init()
}

func main() {	
	Init()
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/", router.Mux)
	http.ListenAndServe(":3000", nil)	
	return
}



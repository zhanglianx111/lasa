package main

import (
	"net/http"
	"router"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/", router.Mux)
	http.ListenAndServe(":3000", nil)
	return
}

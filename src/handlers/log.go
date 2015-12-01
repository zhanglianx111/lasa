package handlers

import(
	"fmt"
	"net/http"
	log "github.com/Sirupsen/logrus"
)

func HandlerSetLogLevel(w http.ResponseWriter, r *http.Request) {
	level := r.URL.Query().Get(":level")
	lvl, err := log.ParseLevel(level)
	if err != nil {
		fmt.Fprintf(w, "not vilid log level:%s", level)
		return
	}
	log.SetLevel(lvl)
	fmt.Fprintf(w, level)
}

func HandlerGetLogLevel(w http.ResponseWriter, r *http.Request) {
	lvl := log.GetLevel()

	fmt.Fprintf(w, lvl.String())
}
package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/martywachocki/gosm/models"
)

var (
	router *mux.Router
)

// Start Starts the webserver
func Start() {
	router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/services", services)
	router.HandleFunc("/services/{serviceID}", service)

	fs := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/").Handler(fs)

	if models.CurrentConfig.Verbose {
		fmt.Println("Starting web UI accessible at http://" + models.CurrentConfig.WebUIHost + ":" + strconv.Itoa(models.CurrentConfig.WebUIPort) + "/")
	}
	err := http.ListenAndServe(models.CurrentConfig.WebUIHost+":"+strconv.Itoa(models.CurrentConfig.WebUIPort), router)
	if err != nil {
		panic(err)
	}
}

package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	port = 8080
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

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		panic(err)
	}
}

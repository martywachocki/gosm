package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
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
	router.HandleFunc("/services/{serviceId}", service)

	fs := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/").Handler(fs)

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		panic(err)
	}
}

func services(writer http.ResponseWriter, request *http.Request) {
	var payload []byte

	switch request.Method {
	case "GET":
		services := models.CurrentConfig.Services
		payload, _ = json.Marshal(services)
	case "POST":
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(payload)
}

func service(writer http.ResponseWriter, request *http.Request) {

}

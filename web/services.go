package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/martywachocki/gosm/models"
)

func services(writer http.ResponseWriter, request *http.Request) {
	var payload []byte

	switch request.Method {
	case "GET":
		payload, _ = json.Marshal(models.CurrentServices)
	case "POST":
		request.ParseForm()
		port := "NULL"
		if request.FormValue("port") != "" {
			port = "'" + request.FormValue("port") + "'"
		}
		models.Database.MustExec("INSERT INTO services (name, protocol, host, port) VALUES('" + request.FormValue("name") + "', '" + request.FormValue("protocol") + "', '" + request.FormValue("host") + "', " + port + ")")
		models.LoadServices()
		payload = []byte("{\"success\":true}")
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(payload)
}

func service(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	var payload []byte
	switch request.Method {
	case "GET":
		var service models.Service
		models.Database.Get(&service, "SELECT * FROM services WHERE id='"+vars["serviceID"]+"'")
		payload, _ = json.Marshal(service)
	case "PUT":
		request.ParseForm()
		port := "NULL"
		if request.FormValue("port") != "" {
			port = "'" + request.FormValue("port") + "'"
		}
		models.Database.MustExec("UPDATE services SET name='" + request.FormValue("name") + "', protocol='" + request.FormValue("protocol") + "', host='" + request.FormValue("host") + "', port=" + port + " WHERE id='" + vars["serviceID"] + "'")
		models.LoadServices()
		payload = []byte("{\"success\":true}")
	case "DELETE":
		models.Database.MustExec("DELETE FROM services WHERE id='" + vars["serviceID"] + "'")
		models.LoadServices()
		payload = []byte("{\"success\":true}")
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(payload)
}

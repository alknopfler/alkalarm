package api

import (
	"github.com/gorilla/mux"
)


func HandlerController()  *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/setup/sensor", HandlerCreateSensor).Methods("POST")
	r.HandleFunc("/setup/sensor", HandlerGetSensors).Methods("GET")
	r.HandleFunc("/setup/sensor/{code}", HandlerDeleteSensor).Methods("DELETE")
	return r
}


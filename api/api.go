package api

import (
	"github.com/gorilla/mux"
)


func HandlerController()  *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/setup/sensor", HandlerCreateSensor).Methods("POST")
	r.HandleFunc("/setup/sensor", HandlerGetSensors).Methods("GET")
	r.HandleFunc("/setup/sensor/{code}", HandlerDeleteSensor).Methods("DELETE")
	r.HandleFunc("/setup/mail", HandlerCreateMail).Methods("POST")
	r.HandleFunc("/setup/mail", HandlerGetMail).Methods("GET")
	r.HandleFunc("/setup/mail/{receptor}", HandlerDeleteMail).Methods("DELETE")
	return r
}


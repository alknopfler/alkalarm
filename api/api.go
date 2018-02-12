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
	r.HandleFunc("/setup/control", HandlerCreateControl).Methods("POST")
	r.HandleFunc("/setup/control", HandlerGetControl).Methods("GET")
	r.HandleFunc("/setup/control/{code}", HandlerDeleteControl).Methods("DELETE")
	r.HandleFunc("/alarm", HandlerGetAlarm).Methods("GET")
	r.HandleFunc("/alarm", HandlerDeleteAlarm).Methods("DELETE")

	r.HandleFunc("/activate/full",HandlerActivateFull).Methods("POST")
	r.HandleFunc("/activate/partial",HandlerActivatePartial).Methods("POST")
	r.HandleFunc("/deactivate",HandlerDeactivate).Methods("POST")
	r.HandleFunc("/status",HandlerAlarmStatus).Methods("GET")
	return r
}


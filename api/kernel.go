package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/kernel"
)

func HandlerActivate(w http.ResponseWriter, r *http.Request) {
	kernel.ListenEvents()
	responseWithJSON(w,http.StatusOK,"Alarm started successfully")
}

func HandlerDeactivate(w http.ResponseWriter, r *http.Request) {
	kernel.State = "stop"
	responseWithJSON(w,http.StatusOK,"Alarm stoped successfully")
}

package api

import (
	"net/http"
	//"github.com/alknopfler/alkalarm/kernel"
	"github.com/alknopfler/alkalarm/states"
	cfg "github.com/alknopfler/alkalarm/config"
	"time"
)

func HandlerActivateFull(w http.ResponseWriter, r *http.Request) {
	if states.Query()!= cfg.STATE_FULL{
		responseWithJSON(w,http.StatusOK,"Alarm full started successfully")
		time.Sleep(10*time.Second)
		states.Update(cfg.STATE_FULL)
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarm is already full activated")
}

func HandlerActivatePartial(w http.ResponseWriter, r *http.Request) {
	if states.Query() != cfg.STATE_PART{
		responseWithJSON(w,http.StatusOK,"Alarm partial started successfully")
		time.Sleep(10*time.Second)
		states.Update(cfg.STATE_PART)
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarm is already partial activated")
}

func HandlerDeactivate(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_FULL || states.Query() == cfg.STATE_PART {
		states.Update(cfg.STATE_INAC)
		responseWithJSON(w, http.StatusOK, "Alarm stopped successfully")
		return
	}
	responseWithJSON(w, http.StatusOK, "Alarm stopped successfully")
}

func HandlerAlarmStatus(w http.ResponseWriter, r *http.Request){
	responseWithJSON(w,http.StatusOK,states.Query())
}
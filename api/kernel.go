package api

import (
	"net/http"
	//"github.com/alknopfler/alkalarm/kernel"
	"github.com/alknopfler/alkalarm/states"
	cfg "github.com/alknopfler/alkalarm/config"
	"time"
)

func HandlerActivateFull(w http.ResponseWriter, r *http.Request) {
	if states.Query()== cfg.STATE_INAC{
		responseWithJSON(w,http.StatusOK,"Alarm started successfully")
		time.Sleep(60 * time.Second)
		states.Update(cfg.STATE_FULL)
		return
	}else if states.Query()== cfg.STATE_PART{
		states.Update(cfg.STATE_FULL)
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarm started successfully")
}

func HandlerActivatePartial(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_INAC{
		states.Update(cfg.STATE_PART)
		responseWithJSON(w,http.StatusOK,"Alarm started successfully")
		return
	}else if states.Query()== cfg.STATE_FULL{
		states.Update(cfg.STATE_PART)
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarm started successfully")
}

func HandlerDeactivate(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_FULL || states.Query() == cfg.STATE_PART {
		states.Update(cfg.STATE_INAC)
		responseWithJSON(w, http.StatusOK, "Alarm stoped successfully")
		return
	}
	responseWithJSON(w, http.StatusOK, "Alarm stoped successfully")
}

func HandlerAlarmStatus(w http.ResponseWriter, r *http.Request){
	responseWithJSON(w,http.StatusOK,states.Query())
}
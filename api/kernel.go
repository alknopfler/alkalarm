package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/kernel"
	"github.com/alknopfler/alkalarm/states"
	cfg "github.com/alknopfler/alkalarm/config"
)

func HandlerActivateFull(w http.ResponseWriter, r *http.Request) {
	if states.Query()== cfg.STATE_INAC{
		states.Update(cfg.STATE_FULL)
		go kernel.ListenEvents()
		responseWithJSON(w,http.StatusOK,"Alarm started successfully")
		return
	}else if states.Query()== cfg.STATE_PART{
		//kernel.State <- "stop"  //primero paro y despues lanzo con full
		states.Update(cfg.STATE_FULL)
		//go kernel.ListenEvents()
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarm started successfully")
}

func HandlerActivatePartial(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_INAC{
		states.Update(cfg.STATE_PART)
		go kernel.ListenEvents()
		responseWithJSON(w,http.StatusOK,"Alarm started successfully")
		return
	}else if states.Query()== cfg.STATE_FULL{
		//kernel.State <- "stop"  //primero paro y despues lanzo con part
		states.Update(cfg.STATE_PART)
		//go kernel.ListenEvents()
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarm started successfully")
}

func HandlerDeactivate(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_FULL || states.Query() == cfg.STATE_PART {
		kernel.State <- "stop"
		states.Update(cfg.STATE_INAC)
		responseWithJSON(w, http.StatusOK, "Alarm stoped successfully")
		return
	}
	responseWithJSON(w, http.StatusOK, "Alarm stoped successfully")
}

func HandlerAlarmStatus(w http.ResponseWriter, r *http.Request){
	responseWithJSON(w,http.StatusOK,states.Query())
}
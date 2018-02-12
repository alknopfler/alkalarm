package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/kernel"
	"github.com/alknopfler/alkalarm/states"
	cfg "github.com/alknopfler/alkalarm/config"
)

func HandlerActivateFull(w http.ResponseWriter, r *http.Request) {
	if states.QueryState()== cfg.STATE_INAC{
		states.UpdateState(cfg.STATE_FULL)
		go kernel.ListenEvents(cfg.STATE_FULL)
		responseWithJSON(w,http.StatusOK,"Alarm started successfully")
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarm started successfully")
}

func HandlerActivatePartial(w http.ResponseWriter, r *http.Request) {
	if states.QueryState() == cfg.STATE_INAC{
		states.UpdateState(cfg.STATE_PART)
		go kernel.ListenEvents(cfg.STATE_PART)
		responseWithJSON(w,http.StatusOK,"Alarm started successfully")
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarm started successfully")
}

func HandlerDeactivate(w http.ResponseWriter, r *http.Request) {
	if states.QueryState() == cfg.STATE_FULL || states.QueryState() == cfg.STATE_PART {
		states.UpdateState(cfg.STATE_INAC)
		kernel.State <- "stop"
		responseWithJSON(w, http.StatusOK, "Alarm stoped successfully")
		return
	}
	responseWithJSON(w, http.StatusOK, "Alarm stoped successfully")
}

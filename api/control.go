package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/control"
	"github.com/gorilla/mux"
	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/states"
)


//HandlerCreateControl function
func HandlerCreateControl(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_INAC {   //must be inactive
		input, err := readControlBodyJson(r)
		if err != nil {
			responseWithError(w, http.StatusBadRequest, err.Error())
		}
		for i := range input {
			err = control.Register(input[i])
			if err != nil {
				responseWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		}
		responseWithJSON(w, http.StatusCreated, "Control Registered successfully")
		return
	}
	responseWithError(w, http.StatusBadGateway, "Alarm state must be inactive")
}

func HandlerScanControl(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_INAC {   //must be inactive
		code,err:=control.ScanControl()
		if err!= nil {
			responseWithError(w,http.StatusInternalServerError,err.Error())
			return
		}
		responseWithJSON(w,http.StatusCreated,code)
		return
	}
	responseWithError(w, http.StatusBadGateway, "Alarm state must be inactive")
}

//HandlerDeleteControl function
func HandlerDeleteControl(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_INAC {   //must be inactive
		codeInput, _ := mux.Vars(r)["code"]
		if ! control.Exists(codeInput){
			responseWithError(w, http.StatusBadGateway, "Control Not Found")
			return
		}
		data:= cfg.Control{Code:codeInput}
		err:=control.Unregister(data.Code)
		if err!=nil{
			responseWithError(w,http.StatusInternalServerError,err.Error())
			return
		}
		responseWithJSON(w,http.StatusOK,"Control Deleted successfully")
		return
	}
	responseWithError(w, http.StatusBadGateway, "Alarm state must be inactive")
}

func HandlerGetControl(w http.ResponseWriter, r *http.Request){
	value,err:=control.QueryAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
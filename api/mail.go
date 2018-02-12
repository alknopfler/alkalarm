package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/mailer"
	"github.com/gorilla/mux"
	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/states"
)


//HandlerCreateMail function
func HandlerCreateMail(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_INAC {   //must be inactive
		input, err := readMailBodyJson(r)
		if err != nil {
			responseWithError(w, http.StatusBadRequest,err.Error())
		}
		for i:=range input{
			err=mailer.Register(input[i])
			if err!= nil {
				responseWithError(w,http.StatusInternalServerError,err.Error())
				return
			}
		}
		responseWithJSON(w,http.StatusCreated,"Mail Registered successfully")
		return
	}
	responseWithError(w, http.StatusBadGateway, "Alarm state must be inactive")
}

//HandlerDeleteMail function
func HandlerDeleteMail(w http.ResponseWriter, r *http.Request) {
	if states.Query() == cfg.STATE_INAC {   //must be inactive
		receptorInput, _ := mux.Vars(r)["receptor"]
		if ! mailer.Exists(receptorInput){
			responseWithError(w, http.StatusBadGateway, "Mail Not Found")
			return
		}
		data:= cfg.Mailer{Receptor:receptorInput}
		err:=mailer.Unregister(data)
		if err!=nil{
			responseWithError(w,http.StatusInternalServerError,err.Error())
			return
		}
		responseWithJSON(w,http.StatusOK,"Mail Deleted successfully")
		return
	}
	responseWithError(w, http.StatusBadGateway, "Alarm state must be inactive")
}

func HandlerGetMail(w http.ResponseWriter, r *http.Request){
	value,err:=mailer.QueryAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
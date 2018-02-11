package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/mailer"
	"github.com/gorilla/mux"
	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/kernel"
)


//HandlerCreateMail function
func HandlerCreateMail(w http.ResponseWriter, r *http.Request) {
	if kernel.GetGlobalState() != cfg.STATE_INAC {   //must be inactive
		input, err := readMailBodyJson(r)
		if err != nil {
			responseWithError(w, http.StatusBadRequest,err.Error())
		}
		for i:=range input{
			err=mailer.RegisterMailer(input[i])
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
	if kernel.GetGlobalState() != cfg.STATE_INAC {   //must be inactive
		receptorInput, _ := mux.Vars(r)["receptor"]
		if ! mailer.MailExists(receptorInput){
			responseWithError(w, http.StatusBadGateway, "Mail Not Found")
			return
		}
		data:= cfg.Mailer{Receptor:receptorInput}
		err:=mailer.UnregisterMailer(data)
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
	value,err:=mailer.QueryMailAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
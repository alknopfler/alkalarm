package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/mailer"
	"github.com/gorilla/mux"
	"errors"
	"github.com/alknopfler/alkalarm/config"
)


//HandlerCreateMail function
func HandlerCreateMail(w http.ResponseWriter, r *http.Request) {
	//TODO la alarma global debe estar desactivada para operar
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
}

//HandlerDeleteMail function
func HandlerDeleteMail(w http.ResponseWriter, r *http.Request) {
	//TODO la alarma global debe estar desactivada para operar
	receptorInput, _ := mux.Vars(r)["receptor"]
	if ! mailer.MailExists(receptorInput){
		responseWithError(w, http.StatusBadGateway, "Mail Not Found")
		return
	}
	data:= config.Mailer{Receptor:receptorInput}
	err:=mailer.UnregisterMailer(data)
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,"Mail Deleted successfully")
}

func HandlerGetMail(w http.ResponseWriter, r *http.Request){
	value,err:=mailer.QueryMailAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/control"
	"github.com/gorilla/mux"
	cfg "github.com/alknopfler/alkalarm/config"
)


//HandlerCreateControl function
func HandlerCreateControl(w http.ResponseWriter, r *http.Request) {
	//TODO la alarma global debe estar desactivada para operar
	input, err := readControlBodyJson(r)
	if err != nil {
		responseWithError(w, http.StatusBadRequest,err.Error())
	}
	for i:=range input{
		err=control.RegisterControl(input[i])
		if err!= nil {
			responseWithError(w,http.StatusInternalServerError,err.Error())
			return
		}
	}
	responseWithJSON(w,http.StatusCreated,"Control Registered successfully")

}

//HandlerDeleteControl function
func HandlerDeleteControl(w http.ResponseWriter, r *http.Request) {
	//TODO la alarma global debe estar desactivada para operar
	codeInput, _ := mux.Vars(r)["code"]
	if ! control.ControlExists(codeInput){
		responseWithError(w, http.StatusBadGateway, "Control Not Found")
		return
	}
	data:= cfg.Control{Code:codeInput}
	err:=control.UnregisterControl(data.Code)
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,"Control Deleted successfully")
}

func HandlerGetControl(w http.ResponseWriter, r *http.Request){
	value,err:=control.QueryControlAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
package api

import (
	"net/http"
	"github.com/gorilla/mux"
	cfg "github.com/alknopfler/alkalarm/config"
)


//HandlerDeleteMail function
func HandlerVerifyPass(w http.ResponseWriter, r *http.Request) {
	receptorInput, _ := mux.Vars(r)["pass"]
	if receptorInput+"\n" != cfg.WEBACCESS_PASS{
		responseWithError(w, http.StatusBadGateway, "Bad Password"+cfg.WEBACCESS_PASS)
		return
	}
	responseWithJSON(w,http.StatusOK,"Good Password")

}


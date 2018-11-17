package api

import (
	"net/http"
)


//HandlerDeleteMail function
func HandlerVerifyPass(w http.ResponseWriter, r *http.Request) {
	//receptorInput, _ := mux.Vars(r)["pass"]
	//if receptorInput != cfg.WEBACCESS_PASS{
	//	responseWithError(w, http.StatusBadGateway, "Bad Password")
	//	return
	//}
	responseWithJSON(w,http.StatusOK,"Good Password")

}


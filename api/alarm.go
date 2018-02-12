package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/alarms"
)



//HandlerDeleteSensor function
func HandlerDeleteAlarm(w http.ResponseWriter, r *http.Request) {

	err:=alarms.Unregister()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarms Deleted successfully")
}

func HandlerGetAlarm(w http.ResponseWriter, r *http.Request){
	value,err:=alarms.QueryAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
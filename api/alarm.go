package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/alarms"
)



//HandlerDeleteSensor function
func HandlerDeleteAlarm(w http.ResponseWriter, r *http.Request) {

	err:=alarms.UnregisterAlarm()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,"Alarms Deleted successfully")
}

func HandlerGetAlarm(w http.ResponseWriter, r *http.Request){
	value,err:=alarms.QueryAlarmAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
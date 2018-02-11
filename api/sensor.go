package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/sensors"
	"github.com/gorilla/mux"
	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/kernel"
)


//HandlerCreateSensor function
func HandlerCreateSensor(w http.ResponseWriter, r *http.Request) {
	if kernel.GetGlobalState() != cfg.STATE_INAC {   //must be inactive
		input, err := readSensorBodyJson(r)
		if err != nil {
			responseWithError(w, http.StatusBadRequest,err.Error())
		}
		for i:=range input{
			err=sensors.RegisterSensor(input[i])
			if err!= nil {
				responseWithError(w,http.StatusInternalServerError,err.Error())
				return
			}
		}
		responseWithJSON(w,http.StatusCreated,"Sensor Registered successfully")
		return
	}
	responseWithError(w, http.StatusBadGateway, "Alarm state must be inactive")
}

//HandlerDeleteSensor function
func HandlerDeleteSensor(w http.ResponseWriter, r *http.Request) {
	if kernel.GetGlobalState() != cfg.STATE_INAC {   //must be inactive
		codeInput, _ := mux.Vars(r)["code"]
		if ! sensors.SensorExists(codeInput){
			responseWithError(w, http.StatusBadGateway, "Sensor Not Found")
			return
		}
		data:= cfg.Sensor{Code:codeInput}
		err:=sensors.UnregisterSensor(data.Code)
		if err!=nil{
			responseWithError(w,http.StatusInternalServerError,err.Error())
			return
		}
		responseWithJSON(w,http.StatusOK,"Sensor Deleted successfully")
		return
	}
	responseWithError(w, http.StatusBadGateway, "Alarm state must be inactive")
}

func HandlerGetSensors(w http.ResponseWriter, r *http.Request){
	value,err:=sensors.QuerySensorsAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
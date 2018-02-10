package api

import (
	"net/http"
	"github.com/alknopfler/alkalarm/sensors"
	"github.com/gorilla/mux"
	cfg "github.com/alknopfler/alkalarm/config"
)


//HandlerCreateSensor function
func HandlerCreateSensor(w http.ResponseWriter, r *http.Request) {
	//TODO la alarma global debe estar desactivada para operar
	input, err := readMailBodyJson(r)
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

}

//HandlerDeleteSensor function
func HandlerDeleteSensor(w http.ResponseWriter, r *http.Request) {
	//TODO la alarma global debe estar desactivada para operar
	codeInput, _ := mux.Vars(r)["code"]
	if ! sensors.SensorExists(codeInput){
		responseWithError(w, http.StatusBadGateway, "Sensor Not Found")
		return
	}
	data:= cfg.Sensor{Code:codeInput}
	err:=sensors.UnregisterSensor(data)
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,"Sensor Deleted successfully")
}

func HandlerGetSensors(w http.ResponseWriter, r *http.Request){
	value,err:=sensors.QuerySensorsAll()
	if err!=nil{
		responseWithError(w,http.StatusInternalServerError,err.Error())
		return
	}
	responseWithJSON(w,http.StatusOK,value)
}
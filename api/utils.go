package api

import (
	"encoding/json"
	"net/http"
	cfg "github.com/alknopfler/alkalarm/config"
	"io/ioutil"
	"log"
)

func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"error": message})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func readSensorBodyJson(r *http.Request)([]cfg.Sensor,error){
	var value []cfg.Sensor
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Println("Error while reading input JSON")
		return value, err
	}

	err = json.Unmarshal(b, &value)
	if err != nil {
		log.Println("Error while unmarshalling input JSON")
		return value, err
	}
	return value, nil
}

func readMailBodyJson(r *http.Request)([]cfg.Mailer,error){
	var value []cfg.Mailer
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Println("Error while reading input JSON")
		return value, err
	}

	err = json.Unmarshal(b, &value)
	if err != nil {
		log.Println("Error while unmarshalling input JSON")
		return value, err
	}
	return value, nil
}

func readControlBodyJson(r *http.Request)([]cfg.Control,error){
	var value []cfg.Control
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Println("Error while reading input JSON")
		return value, err
	}

	err = json.Unmarshal(b, &value)
	if err != nil {
		log.Println("Error while unmarshalling input JSON")
		return value, err
	}
	return value, nil
}

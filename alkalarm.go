package main

import (
	_ "github.com/mattn/go-sqlite3"
	"os"
	"github.com/alknopfler/alkalarm/database"
	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/api"
	"github.com/gorilla/handlers"
	"net/http"
	"github.com/alknopfler/alkalarm/kernel"
	"log/syslog"
	"log"
	"github.com/gorilla/mux"
)


func init(){
	needCreation:=true
 	logging, e := syslog.New(syslog.LOG_NOTICE, "alkalarm")
	if e == nil {
		log.SetOutput(logging)
	}

	log.Print("Validating the database, and other params...Could take some minutes...")
	//First Time to execute needs create database and scheme
	if _, err := os.Stat(cfg.DB_NAME); err == nil {
		log.Println("ya existe la base de datos")
		needCreation= false
	}
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB First Time (init)")
		os.Exit(2)
	}
	defer db.Close()
	err=database.CreateSchemas(db)
	if err!=nil{
		log.Println("Error creating the schema the first time (init)")
		os.Exit(3)
	}
	if needCreation {
		err=database.Operate(db,cfg.GLOBAL_STATE_INSERT,cfg.STATE_INAC)
		if err!=nil{
			log.Println("Error activating the first time (init)")
			os.Exit(3)
		}
		//err=database.Operate(db,cfg.ADMIN_INSERT,cfg.WEBACCESS_PASS)
		//if err!=nil{
		//	log.Println("Error creating the first pass (init)")
		//	os.Exit(3)
		//}
	}
	log.Println("Success...Starting the program")
}

func main() {
	go kernel.ListenEvents()  //lanzo el primero por si la activo con el mando en lugar de con la api

	r := mux.NewRouter()
	r.HandleFunc("/setup/sensor", api.HandlerCreateSensor).Methods("POST")
	r.HandleFunc("/scan/sensor", api.HandlerScanSensor).Methods("GET")
	r.HandleFunc("/sensors", api.HandlerGetSensors).Methods("GET")
	r.HandleFunc("/setup/sensor/{code}", api.HandlerDeleteSensor).Methods("DELETE")
	r.HandleFunc("/setup/mail", api.HandlerCreateMail).Methods("POST")
	r.HandleFunc("/mails", api.HandlerGetMail).Methods("GET")
	r.HandleFunc("/setup/mail/{receptor}", api.HandlerDeleteMail).Methods("DELETE")
	r.HandleFunc("/setup/control", api.HandlerCreateControl).Methods("POST")
	r.HandleFunc("/scan/control", api.HandlerScanControl).Methods("GET")
	r.HandleFunc("/controls", api.HandlerGetControl).Methods("GET")
	r.HandleFunc("/setup/control/{code}", api.HandlerDeleteControl).Methods("DELETE")
	r.HandleFunc("/alarm", api.HandlerGetAlarm).Methods("GET")
	r.HandleFunc("/alarm", api.HandlerDeleteAlarm).Methods("DELETE")

	r.HandleFunc("/activate/full",api.HandlerActivateFull).Methods("POST")
	r.HandleFunc("/activate/partial",api.HandlerActivatePartial).Methods("POST")
	r.HandleFunc("/deactivate",api.HandlerDeactivate).Methods("POST")
	r.HandleFunc("/status",api.HandlerAlarmStatus).Methods("GET")
	r.HandleFunc("/admin/{pass}",api.HandlerVerifyPass).Methods("GET")

	corsObj:=handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	err := http.ListenAndServe(cfg.SERVER_API_PORT, handlers.CORS(corsObj,headersOk,methodsOk)(r))
	if err != nil {
		log.Println("Error listening api server...")
	}
}

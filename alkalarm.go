package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"github.com/alknopfler/alkalarm/database"
	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/api"
	"net/http"
	"github.com/alknopfler/alkalarm/kernel"
	"log/syslog"
	"log"
)


func init(){
	logging, e := syslog.New(syslog.LOG_NOTICE, "alkalarm")
	if e == nil {
		log.SetOutput(logging)
	}

	log.Print("Validating the database, and other params...Could take some minutes...")
	//First Time to execute needs create database and scheme
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
	if _, err := os.Stat(cfg.DB_NAME); err == nil {
		err=database.Operate(db,cfg.GLOBAL_STATE_UPDATE,cfg.STATE_INAC)
		if err!=nil{
			log.Println("Error activating the first time (init)")
			os.Exit(3)
		}
	}else{
		err=database.Operate(db,cfg.GLOBAL_STATE_INSERT,cfg.STATE_INAC)
		if err!=nil{
			log.Println("Error activating the first time (init)")
			os.Exit(3)
		}
	}

	log.Println("Success...Starting the program")
}

func main() {
	go kernel.ListenEvents()  //lanzo el primero por si la activo con el mando en lugar de con la api
	err := http.ListenAndServe(cfg.SERVER_API_PORT, api.HandlerController())
	if err != nil {
		log.Println("Error listening api server...")
	}
}
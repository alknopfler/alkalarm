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
)



func init(){

	fmt.Println("Validating the database, and other params...Could take some minutes...")
	//First Time to execute needs create database and scheme
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB First Time (init):",err)
		os.Exit(2)
	}
	defer db.Close()
	//TODO revisar si es necesario crear esquemas o simplemente arrancar
	err=database.CreateSchemas(db)
	if err!=nil{
		fmt.Println("Error creating the schema the first time (init):", err)
		os.Exit(3)
	}
	err=database.Operate(db,cfg.GLOBAL_STATE_INSERT,cfg.STATE_INAC)
	if err!=nil{
		fmt.Println("Error activating the first time (init):", err)
		os.Exit(3)
	}
	fmt.Println("Success...Starting the program")
}

func main() {
	//go kernel.ListenEvents()  //lanzo el primero por si la activo con el mando en lugar de con la api
	err := http.ListenAndServe(cfg.SERVER_API_PORT, api.HandlerController())
	if err != nil {
		fmt.Println("Error listening api server...")
	}
}

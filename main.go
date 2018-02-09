package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/alknopfler/alkalarm/database"
	"github.com/alknopfler/alkalarm/sensors"
)



func init(){
	fmt.Println("Validating the database, and other params...Could take some minutes...")
	//First Time to execute needs create database and scheme
	db,_ := database.InitDB()
	defer db.Close()
	database.CreateSchemas(db)
	fmt.Println("Success...Starting the program")
}

func main() {
	sensors.RegisterSensor()
}

package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/alknopfler/alkalarm/database"
	"github.com/alknopfler/alkalarm/sensors"
	"os"
)



func init(){
	fmt.Println("Validating the database, and other params...Could take some minutes...")
	//First Time to execute needs create database and scheme
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB")
		os.Exit(2)
	}
	defer db.Close()
	database.CreateSchemas(db)
	fmt.Println("Success...Starting the program")
}

func main() {
	sensors.RegisterSensor()
}

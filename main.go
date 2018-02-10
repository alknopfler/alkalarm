package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"github.com/alknopfler/alkalarm/database"
	//"github.com/alknopfler/alkalarm/api"
	//"net/http"
	"github.com/alknopfler/alkalarm/mailer"
	cfg "github.com/alknopfler/alkalarm/config"
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
	err=database.CreateSchemas(db)
	if err!=nil{
		fmt.Println("Error creating the schema the first time (init):", err)
		os.Exit(3)
	}
	fmt.Println("Success...Starting the program")
}

func main() {
	/*err := http.ListenAndServe(":8080", api.HandlerController())
	if err != nil {
		fmt.Println("Error listening server in :8080")
	}*/
	data:=cfg.Mailer{Receptor: "alknopfler@hotmail.com"}
	mailer.RegisterMailer(data)
	data = cfg.Mailer{Receptor: "begoclavero@gmail.com"}
	mailer.RegisterMailer(data)
	mailer.SendMail("ventana Salon")
}

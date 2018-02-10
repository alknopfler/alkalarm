package mailer

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"fmt"
	"github.com/alknopfler/alkalarm/database"
)

func RegisterMailer(data cfg.Mailer) error{
	//First Time to execute needs create database and scheme
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Register Mailer")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.MAIL_INSERT,data.Emisor,data.Receptor,data.Subject,data.Text,data.Smtp_address,data.Smtp_port,data.Smtp_user,data.Smtp_pass,data.Smtp_security)
	if err!=nil{
		fmt.Println("Error inserting mailer in db")
		return err
	}
	fmt.Println("Success...Mail registered successfully")
	return nil
}

func UnregisterMailer(data cfg.Mailer) error{
	//First Time to execute needs create database and scheme
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Register Mailer")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.MAIL_DELETE,data.Receptor)
	if err!=nil{
		fmt.Println("Error inserting mailer in db")
		return err
	}
	fmt.Println("Success...Mail registered successfully")
	return nil
}

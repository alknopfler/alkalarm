package mailer

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"net/smtp"
	"log"
	"github.com/alknopfler/alkalarm/database"

	"strings"
)
func Register(data cfg.Mailer) error{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Register Mailer")
		return err
	}
	defer db.Close()
	err=database.Operate(db,cfg.MAIL_INSERT,data.Receptor)
	if err!=nil{
		log.Println("Error inserting mailer in db")
		return err
	}
	log.Println("Success...Mail registered successfully")
	return nil
}

func Unregister(data cfg.Mailer) error{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Register Mailer")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.MAIL_DELETE,data.Receptor)
	if err!=nil{
		log.Println("Error inserting mailer in db")
		return err
	}
	log.Println("Success...Mail unregistered successfully")
	return nil
}

func SendMail(typeof,zona string){
	//first of all get all the mails to send the emails
	list,err:=QueryAll()
	if err != nil {
		log.Println("Error retrieving the mails to send")
		return
	}

	msg := "From: " + cfg.FROM + "\n" +
		"To: " + list[0].Receptor + "\n" +
		"Subject: ALARMA CASA - Sensor de "+strings.ToUpper(typeof)+" de la zona: "+strings.ToUpper(zona)+"\n\n" +
		"Ha saltado el sensor de "+strings.ToUpper(typeof)+" de la zona : " + strings.ToUpper(zona)

	err = smtp.SendMail(cfg.SMTP_SERVER+":"+cfg.SMTP_PORT,
		smtp.PlainAuth("", cfg.FROM, cfg.SMTP_PASS, cfg.SMTP_SERVER),
		cfg.FROM, getArrayMailTo(list), []byte(msg))

	if err != nil {
		log.Println("smtp error:", err)
		return
	}
}

func getArrayMailTo(list []cfg.Mailer)[]string{
	var result [len(list)]string
	for i := range list {
		result[i]=list[i].Receptor
	}
	return result
}

func QueryAll() ([]cfg.Mailer,error){
	var result []cfg.Mailer
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Query Sensor")
		return result,err
	}
	defer db.Close()
	rows, err := db.Query(cfg.MAIL_QUERY_ALL)
	if err != nil { return result,err }
	defer rows.Close()
	var item cfg.Mailer
	for rows.Next() {
		item.Receptor = ""
		err2 := rows.Scan(&item.Receptor)
		if err2 != nil { return nil,err }
		result = append(result, item)
	}
	return result, nil
}

func Exists(receptor string) bool{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Mail Exists")
		return false
	}
	defer db.Close()
	rows, err := db.Query(cfg.MAIL_QUERY_RECEPTOR,receptor)
	defer rows.Close()
	if rows.Next(){
		return true
	}
	return false
}

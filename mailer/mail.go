package mailer

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"net/smtp"
	"log"
)

func SendMail(zona string,){

	msg := "From: " + cfg.FROM + "\n" +
		"To: " + cfg.LIST_TO_MAIL[0] + "\n" +
		"Subject: ALARMA CASA - Sensor zona: "+zona+"\n\n" +
		"Ha saltado la alarma de la casa del sensor: " + zona

	err := smtp.SendMail(cfg.SMTP_SERVER+":"+cfg.SMTP_PORT,
		smtp.PlainAuth("", cfg.FROM, cfg.SMTP_PASS, cfg.SMTP_SERVER),
		cfg.FROM, cfg.LIST_TO_MAIL, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

}
package database

const (
	//// INSERTS ////
	SENSOR_INSERT = "INSERT INTO sensors(code,type,zone) values(?,?,?)"
	MAIL_INSERT = "INSERT INTO mailer(emisor,receptor,subject,text,smtp_address,smtp_port,smtp_user,smtp_pass,smtp_security) values (?,?,?,?,?,?,?,?,?)"
	ALARM_INSERT = "INSERT INTO alarms(date,sensor) values(?,?)"
	CONTROL_INSERT = "INSERT INTO control(code,type) values(?,?)"
	GLOBAL_STATE_INSERT= "INSERT INTO global_state(state) values(?)"

	//// SELECTS ////

	//// DELETES ////

	//// UPDATES ////
)

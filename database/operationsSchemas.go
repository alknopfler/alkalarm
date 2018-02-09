package database

/*
		sensors: insert | select | delete       | (not make sense update)
		mail:    insert | select | delete       | (not make sense update)
		alarm:   insert | select | delete (all) | (not make sense update)
		control: insert | select | delete       | (not make sense update)
		global:  insert | select | delete (all) | update


 */

const (
	//// INSERTS ////
		SENSOR_INSERT = "INSERT INTO sensors(code,typeOf,zone) values(?,?,?)"
		MAIL_INSERT = "INSERT INTO mailer(emisor,receptor,subject,text,smtp_address,smtp_port,smtp_user,smtp_pass,smtp_security) values (?,?,?,?,?,?,?,?,?)"
		ALARM_INSERT = "INSERT INTO alarms(date,sensor) values(?,?)"
		CONTROL_INSERT = "INSERT INTO control(code,description,typeOf) values(?,?,?)"
		GLOBAL_STATE_INSERT= "INSERT INTO global_state(id,state) values(1,?)"

	//// DELETES ////
		SENSOR_DELETE = "DELETE FROM sensors WHERE code=?"
		MAIL_DELETE = "DELETE FROM mailer WHERE receptor=?"
		ALARM_DELETE = "DELETE FROM alarms"   //just to clear alarm historic
		CONTROL_DELETE = "DELETE FROM control WHERE code=?"
		GLOBAL_STATE_DELETE = "DELETE FROM global_state"

	//// UPDATES ////
		GLOBAL_STATE_UPDATE = "UPDATE global_state SET state=? WHERE id=1"


	//// QUERIES ////
		SENSOR_QUERY = "SELECT * FROM sensors"

)

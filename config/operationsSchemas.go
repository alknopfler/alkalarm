package config

/*
		sensors: insert | select | delete       | (not make sense update)
		alarm:   insert | select | delete (all) | (not make sense update)
		mail:    insert | select | delete (all) | (not make sense update)
		control: insert | select | delete       | (not make sense update)
		global:  insert | select | delete (all) | update
 */

const (
	//// INSERTS ////
		SENSOR_INSERT = "INSERT INTO sensors(code,typeOf,zone) values(?,?,?)"
		MAIL_INSERT = "INSERT INTO mailer(receptor) values (?)"
		ALARM_INSERT = "INSERT INTO alarms(date,sensor) values(?,?)"
		CONTROL_INSERT = "INSERT INTO control(code,description,typeOf) values(?,?,?)"
		GLOBAL_STATE_INSERT= "INSERT INTO global_state(id,gstate) values(1,?)"
		ADMIN_INSERT="INSERT INTO admin(pass) values(?)"

	//// DELETES ////
		SENSOR_DELETE = "DELETE FROM sensors WHERE code=?"
		MAIL_DELETE = "DELETE FROM mailer WHERE receptor=?"
		ALARM_DELETE = "DELETE FROM alarms"   //just to clear alarm historic
		CONTROL_DELETE = "DELETE FROM control WHERE code=?"

	//// UPDATES ////
		GLOBAL_STATE_UPDATE = "UPDATE global_state SET gstate=? WHERE id=1"


	//// QUERIES ////
		SENSOR_QUERY_ALL= "SELECT * FROM sensors"
		SENSOR_QUERY_CODE = "SELECT * FROM sensors WHERE code=?"
		MAIL_QUERY_ALL = "SELECT * FROM mailer"
		MAIL_QUERY_RECEPTOR = "SELECT * FROM mailer WHERE receptor=?"
		CONTROL_QUERY_ALL= "SELECT * FROM control"
		CONTROL_QUERY_CODE = "SELECT * FROM control WHERE code=?"
		ALARM_QUERY_ALL= "SELECT * FROM alarms"
		GLOBAL_STATE_QUERY="SELECT * FROM global_state WHERE id=1"

)

package config


const (
	//// TABLES ////
	SENSORS_TABLE = "CREATE TABLE IF NOT EXISTS sensors (" +
		"code TEXT PRIMARY KEY," +
		"typeOf TEXT CHECK( typeOf IN ('presence','aperture','other') )," +
		"zone TEXT);"
	ALARM_HISTORY_TABLE = "CREATE TABLE IF NOT EXISTS alarms (" +
		"date TEXT PRIMARY KEY," +
		"sensor TEXT);"
	MAIL_TABLE = "CREATE TABLE IF NOT EXISTS mailer (" +
		"receptor TEXT PRIMARY KEY)"
	CONTROL_CODES_TABLE = "CREATE TABLE IF NOT EXISTS control (" +
		"code TEXT PRIMARY KEY," +
		"description TEXT," +
		"typeOf TEXT CHECK( typeOf IN ('"+STATE_INAC+"','"+STATE_FULL+"','"+STATE_PART+"','"+STATE_SOS+"') ) NOT NULL DEFAULT '');"
	GLOBAL_STATE_TABLE = "CREATE TABLE IF NOT EXISTS global_state (" +
		"id TEXT PRIMARY KEY,"+
		"gstate TEXT CHECK( gstate IN ('"+STATE_FULL+"','"+STATE_PART+"','"+STATE_INAC+"') ) NOT NULL DEFAULT '');"
	ADMIN_TABLE = "CREATE TABLE IF NOT EXISTS admin (" +
		"pass TEXT PRIMARY KEY;"

)


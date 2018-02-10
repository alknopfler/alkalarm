package config


const (
	//// TABLES ////
	SENSORS_TABLE = "CREATE TABLE IF NOT EXISTS sensors (" +
		"code TEXT PRIMARY KEY," +
		"typeOf TEXT CHECK( typeOf IN ('presence','aperture','other') )," +
		"zone TEXT);"
	MAILER_TABLE = "CREATE TABLE IF NOT EXISTS mailer (" +
		"emisor TEXT," +
		"receptor TEXT," +
		"subject TEXT," +
		"text TEXT," +
		"smtp_address TEXT," +
		"smtp_port TEXT," +
		"smtp_user TEXT," +
		"smtp_pass TEXT," +
		"smtp_security TEXT);"
	ALARM_HISTORY_TABLE = "CREATE TABLE IF NOT EXISTS alarms (" +
		"date TEXT," +
		"sensor TEXT);"
	CONTROL_CODES_TABLE = "CREATE TABLE IF NOT EXISTS control (" +
		"code TEXT PRIMARY KEY," +
		"description TEXT," +
		"typeOf TEXT CHECK( typeOf IN ('"+STATE_INAC+"','"+STATE_FULL+"','"+STATE_PART+"','"+STATE_SOS+"') ) NOT NULL DEFAULT '');"

	GLOBAL_STATE_TABLE = "CREATE TABLE IF NOT EXISTS global_state (" +
		"id TEXT PRIMARY KEY,"+
		"state TEXT CHECK( state IN ('"+STATE_FULL+"','"+STATE_PART+"','"+STATE_INAC+"') ) NOT NULL DEFAULT '');"
)


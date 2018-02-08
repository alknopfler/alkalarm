package database

import "github.com/alknopfler/alkalarm/config"

const (
	SENSORS_TABLE = "CREATE TABLE IF NOT EXISTS sensors (" +
		"code INTEGER PRIMARY KEY," +
		"type TEXT CHECK( type IN ('presence','aperture','other') ), zone TEXT)"
	MAILER_TABLE = "CREATE TABLE IF NOT EXISTS mailer (" +
		"emisor TEXT," +
		"receptor TEXT," +
		"subject TEXT," +
		"text TEXT," +
		"smtp_address TEXT," +
		"smtp_port TEXT," +
		"smtp_user TEXT," +
		"smtp_pass TEXT," +
		"smtp_security TEXT)"
	ALARM_HISTORY_TABLE = "CREATE TABLE IF NOT EXISTS alarms (" +
		"date TEXT," +
		"sensor INTEGER)"
	CONTROL_CODES_TABLE = "CREATE TABLE IF NOT EXISTS control (" +
		"code INTEGER PRIMARY KEY," +
		"type TEXT CHECK( type IN ('"+config.STATE_INAC+"','"+config.STATE_FULL+"','"+config.STATE_PART+"','"+config.STATE_SOS+"') ) NOT NULL DEFAULT '')"

	GLOBAL_STATE = "CREATE TABLE IF NOT EXISTS global_state (" +
		"state TEXT CHECK( state IN ('"+config.STATE_FULL+"','"+config.STATE_PART+"','"+config.STATE_INAC+"') ) NOT NULL DEFAULT '')"
)



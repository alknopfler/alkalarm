package database

import(
	_ "github.com/mattn/go-sqlite3"

	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestInitDB(t *testing.T) {
	db,err := InitDB()
	defer db.Close()
	assert.NoError(t,err)
}

func TestCreateSchemas(t *testing.T) {
	db,_ := InitDB()
	defer db.Close()
	err:=CreateSchemas(db)
	assert.NoError(t,err)
}


func TestOperateWithItem(t *testing.T) {
	db,_ := InitDB()
	defer db.Close()
	//prueba de insert y delete sobre sensor
	OperateWithItem(db, SENSOR_INSERT, "sensorTest","presence","pasilloTest")
	err:=OperateWithItem(db, SENSOR_DELETE, "sensorTest")
	assert.NoError(t,err)

	//prueba de insert y delete sobre mailer
	OperateWithItem(db, MAIL_INSERT, "emisorTest","receptorTest","subjectTest","textTest","smtp","smtp2","smtp3","smtp4","smtp5")
	err=OperateWithItem(db, MAIL_DELETE, "receptorTest")
	assert.NoError(t,err)

	//prueba de insert y delete sobre alarm
	OperateWithItem(db, ALARM_INSERT, "fechaTest","sensorTest")
	err=OperateWithItem(db, ALARM_DELETE)
	assert.NoError(t,err)

	//prueba de insert y delete sobre control
	OperateWithItem(db, CONTROL_INSERT, "codeTest","presence","full")
	err=OperateWithItem(db, CONTROL_DELETE, "codeTest")
	assert.NoError(t,err)

	//prueba de insert , update and delete sobre sensor
	OperateWithItem(db, GLOBAL_STATE_INSERT, "1","full")
	OperateWithItem(db, GLOBAL_STATE_UPDATE, "1","partial")
	err=OperateWithItem(db, GLOBAL_STATE_DELETE)
	assert.NoError(t,err)


	os.Remove("data.db")
}



package database

import(
	_ "github.com/mattn/go-sqlite3"

	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	cfg "github.com/alknopfler/alkalarm/config"
)

func TestInitDB(t *testing.T) {
	db,err := InitDB("./test.db")
	defer db.Close()
	assert.NoError(t,err)
}

func TestCreateSchemas(t *testing.T) {
	db,_ := InitDB("./test.db")
	defer db.Close()
	err:=CreateSchemas(db)
	assert.NoError(t,err)
}


func TestOperateWithItem(t *testing.T) {
	db,_ := InitDB("./test.db")
	defer db.Close()
	//prueba de insert y delete sobre sensor
	Operate(db, cfg.SENSOR_INSERT, "sensorTest","presence","pasilloTest")
	err:=Operate(db, cfg.SENSOR_DELETE, "sensorTest")
	assert.NoError(t,err)

	//prueba de insert y delete sobre alarm
	Operate(db, cfg.ALARM_INSERT, "fechaTest","sensorTest")
	err=Operate(db, cfg.ALARM_DELETE)
	assert.NoError(t,err)

	//prueba de insert y delete sobre control
	Operate(db, cfg.CONTROL_INSERT, "codeTest","presence","full")
	err=Operate(db, cfg.CONTROL_DELETE, "codeTest")
	assert.NoError(t,err)


	os.Remove("./test.db")
}



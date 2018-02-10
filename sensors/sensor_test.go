package sensors

import (
	"testing"
	"github.com/stretchr/testify/assert"
	cfg "github.com/alknopfler/alkalarm/config"
)

func TestRegisterSensor(t *testing.T) {
	data:=cfg.Sensor{
		Code:"testCode",
		TypeOf: "presence",
		Zone: "ventanaTest"}
	err:=RegisterSensor(data)
	assert.NoError(t,err)
}

func TestQuerySensors(t *testing.T) {
	code,err:=QuerySensors("testCode")
	assert.NoError(t,err)
	assert.Equal(t,code,cfg.Sensor{Code:"testCode", TypeOf:"presence", Zone:"ventanaTest"})
}

func TestQuerySensorsAll(t *testing.T) {
	_,err:=QuerySensorsAll()
	assert.NoError(t,err)
}

func TestSensorExists(t *testing.T) {
	res:=SensorExists("testCode")
	assert.True(t,res)
	res2:=SensorExists("codeeeeeefail")
	assert.False(t,res2)
}

func TestUnregisterSensor(t *testing.T) {
	err:=UnregisterSensor("testCode")
	assert.NoError(t,err)
}
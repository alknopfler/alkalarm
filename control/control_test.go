package control

import (
	"testing"
	"github.com/stretchr/testify/assert"
	cfg "github.com/alknopfler/alkalarm/config"
)

func TestRegisterSensor(t *testing.T) {
	data:=cfg.Control{
		Code:"testCode",
		Description: "mando1",
		TypeOf: "full"}
	err:=RegisterControl(data)
	assert.NoError(t,err)
}

func TestQuerySensors(t *testing.T) {
	code,err:=QueryControl("testCode")
	assert.NoError(t,err)
	assert.Equal(t,code,cfg.Control{Code:"testCode", Description:"mando1", TypeOf:"full"})
}

func TestQuerySensorsAll(t *testing.T) {
	_,err:=QueryControlAll()
	assert.NoError(t,err)
}

func TestSensorExists(t *testing.T) {
	res:=ControlExists("testCode")
	assert.True(t,res)
	res2:=ControlExists("codeeeeeefail")
	assert.False(t,res2)
}

func TestUnregisterSensor(t *testing.T) {
	err:=UnregisterControl("testCode")
	assert.NoError(t,err)
}
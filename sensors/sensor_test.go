package sensors

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/alknopfler/alkalarm/database"
)

func TestRegisterSensor(t *testing.T) {
	db,err := database.InitDB("./test.db")
	defer db.Close()
	assert.NoError(t,err)

	//RegisterSensor()
}
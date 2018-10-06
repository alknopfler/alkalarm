package kernel

import (
	"github.com/alknopfler/alkalarm/sensors"
	"github.com/alknopfler/alkalarm/mailer"
	cfg "github.com/alknopfler/alkalarm/config"

	"time"
	"github.com/alknopfler/alkalarm/alarms"
	"sync"
)


func Notificate(evento string){
	var wg *sync.WaitGroup
	//TODO GPIO con alarma sirena
	sensor,_ := sensors.Query(evento)
	data := cfg.Alarm{
		Date: time.Now().Format("Mon Jan _2 15:04:05 2006"),
		Sensor: sensor.Zone+"  "+sensor.TypeOf+"  "+sensor.Code,
	}
	wg.Add(2)

	go mailer.SendMail(sensor.TypeOf,sensor.Zone, wg)  //envio y me desentiendo

	go alarms.Register(data, wg)  //registro alarma y me desentiendo

	wg.Wait()

}
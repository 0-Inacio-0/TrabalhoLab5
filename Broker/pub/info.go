package pub

import (
	"TrabalhoLab5/Broker/config"
	"time"
)

// info <data-hora: char[16] "dd/mm/aaaa hh:mm">
func Info(duration time.Time) {
	config.Pub("info", []byte("<data-hora: "+duration.Format("02/01/2006 15:04")+">"))
}

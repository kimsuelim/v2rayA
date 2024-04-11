package cloud

import (
	"fmt"
	"github.com/v2rayA/v2rayA/conf"
	"time"
)

var (
	FailCreate = fmt.Errorf("failed to create")
	FailGet    = fmt.Errorf("failed to get")
)

func TickDuration() time.Duration {
	var duration time.Duration

	if conf.IsDebug() {
		duration = time.Second * 10
	} else {
		duration = time.Minute
	}

	return duration
}

func GetApiHost() string {
	var url string

	if conf.IsDebug() {
		url = "http://host.docker.internal:8080"
	} else {
		url = "https://imc-api.mooo.com"
	}

	return url
}

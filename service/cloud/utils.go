package cloud

import "github.com/v2rayA/v2rayA/conf"

func GetApiHost() string {
	var url string

	if conf.IsDebug() {
		url = "http://host.docker.internal:8080"
	} else {
		url = "https://imc-api.mooo.com"
	}

	return url
}

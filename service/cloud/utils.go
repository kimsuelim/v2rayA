package cloud

import (
	"fmt"
	"time"

	"github.com/v2rayA/v2rayA/conf"
	"github.com/v2rayA/v2rayA/core/v2ray/where"
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
	return conf.ApiHost
}

func V2rayServiceVersion() string {
	variant, version, err := where.GetV2rayServiceVersion()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s %s", variant, version)
}

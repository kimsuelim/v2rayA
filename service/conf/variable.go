package conf

import (
	"time"
)

var (
	BasicAuthUsername        = "my_username"
	BasicAuthPassword        = "my_password"
	Version                  = "debug"
	FoundNew                 = false
	RemoteVersion            = ""
	TickerUpdateGFWList      *time.Ticker
	TickerUpdateSubscription *time.Ticker
)

func IsDebug() bool {
	return Version == "debug"
}

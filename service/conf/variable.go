package conf

import (
	"time"
)

var (
	AdminUsername            = "admin@admin"
	AdminPassword            = "admin"
	ApiHost                  = "http://host.docker.internal:3000"
	SentryDSN                = "https://5d1d8b27004195f9a0d9f996165b9a96@o4507083982700544.ingest.us.sentry.io/4507084478939136"
	Version                  = "debug"
	FoundNew                 = false
	RemoteVersion            = ""
	TickerUpdateGFWList      *time.Ticker
	TickerUpdateSubscription *time.Ticker
)

func IsDebug() bool {
	return Version == "debug"
}

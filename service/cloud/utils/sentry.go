package utils

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/v2rayA/v2rayA/conf"
	"github.com/v2rayA/v2rayA/db/configure"
)

func AttachCustomContextToSentryEvent() {
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		gs := configure.GetServers()
		for i, s := range gs {
			scope.SetContext(fmt.Sprintf("Server %v", i+1), map[string]interface{}{
				"hostname": s.ServerObj.GetHostname(),
				"proto":    s.ServerObj.ProtoToShow(),
				"port":     s.ServerObj.GetPort(),
			})
		}

		gcs := configure.GetConnectedServers().Get()
		for i, t := range gcs {
			s, _ := t.LocateServerRaw()
			scope.SetContext(fmt.Sprintf("Connected Server %v", i+1), map[string]interface{}{
				"hostname": s.ServerObj.GetHostname(),
				"proto":    s.ServerObj.ProtoToShow(),
				"port":     s.ServerObj.GetPort(),
			})
		}

		scope.SetContext("App", map[string]interface{}{
			"uuid":    configure.GetUUID(),
			"version": conf.Version,
		})
	})
}

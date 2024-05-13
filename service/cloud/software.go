package cloud

import (
	"runtime"

	"github.com/v2rayA/v2rayA/conf"
)

type SoftwareInfo struct {
	OS           string `json:"os"`
	Arch         string `json:"arch"`
	GoVersion    string `json:"goVersion"`
	SwVersion    string `json:"swVersion"`
	V2rayVersion string `json:"v2rayVersion"`
}

func GetSoftwareInfo() SoftwareInfo {
	sw := SoftwareInfo{
		OS:           runtime.GOOS,
		Arch:         runtime.GOARCH,
		GoVersion:    runtime.Version(),
		SwVersion:    conf.Version,
		V2rayVersion: V2rayServiceVersion(),
	}

	return sw
}

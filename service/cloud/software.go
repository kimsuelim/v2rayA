package cloud

import (
	"github.com/v2rayA/v2rayA/conf"
	"runtime"
)

type SoftwareInfo struct {
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	GoVersion string `json:"goVersion"`
	SwVersion string `json:"swVersion"`
}

func GetSoftwareInfo() SoftwareInfo {
	sw := SoftwareInfo{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		SwVersion: conf.Version,
	}

	return sw
}

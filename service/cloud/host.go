package cloud

import (
	"time"

	"github.com/elastic/go-sysinfo"
)

type SysInfo struct {
	Architecture       string    `json:"architecture"`
	NativeArchitecture string    `json:"native_architecture"`
	BootTime           time.Time `json:"boot_time"`
	Containerized      bool      `json:"containerized"`
	Hostname           string    `json:"hostname"`
	IPs                []string  `json:"ip"`
	KernelVersion      string    `json:"kernel_version"`
	MACs               []string  `json:"mac"`
	Os                 struct {
		Type     string `json:"type"`
		Family   string `json:"family"`
		Platform string `json:"platform"`
		Name     string `json:"name"`
		Version  string `json:"version"`
		Major    int    `json:"major"`
		Minor    int    `json:"minor"`
		Patch    int    `json:"patch"`
	} `json:"os"`
	Timezone          string `json:"timezone"`
	TimezoneOffsetSec int    `json:"timezone_offset_sec"`
	UniqueID          string `json:"unique_id"`
}

type HostInfo struct {
	Hostname          string    `json:"hostname"`
	BootTime          time.Time `json:"bootTime"`
	Os                string    `json:"os"`
	Platform          string    `json:"platform"`
	PlatformFamily    string    `json:"platformFamily"`
	PlatformVersion   string    `json:"platformVersion"`
	KernelVersion     string    `json:"kernelVersion"`
	KernelArch        string    `json:"kernelArch"`
	Timezone          string    `json:"timezone"`
	TimezoneOffsetSec int       `json:"timezoneOffsetSec"`
}

func GetHostInfo() HostInfo {
	host, err := sysinfo.Host()
	if err != nil {
		return HostInfo{}
	}

	var hostInfo = HostInfo{
		Hostname:          host.Info().Hostname,
		BootTime:          host.Info().BootTime,
		Os:                host.Info().OS.Type,
		Platform:          host.Info().OS.Platform,
		PlatformFamily:    host.Info().OS.Family,
		PlatformVersion:   host.Info().OS.Version,
		KernelVersion:     host.Info().KernelVersion,
		KernelArch:        host.Info().Architecture,
		Timezone:          host.Info().Timezone,
		TimezoneOffsetSec: host.Info().TimezoneOffsetSec,
	}

	return hostInfo
}

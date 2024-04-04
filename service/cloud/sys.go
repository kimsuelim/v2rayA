package cloud

import (
	"encoding/json"
	"github.com/elastic/go-sysinfo"
	"log"
	"time"
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

func GetSysInfo() string {
	host, err := sysinfo.Host()
	if err != nil {
		return ""
	}

	reqBody, err := json.Marshal(host.Info())
	if err != nil {
		panic(err)
	}

	log.Printf("SysInfo: %v", string(reqBody))

	return string(reqBody)
}

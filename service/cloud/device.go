package cloud

import "github.com/v2rayA/v2rayA/db/configure"

type DeviceInfo struct {
	UUID         string        `json:"uuid"`
	NetworkInfo  []NetworkInfo `json:"networkInfo"`
	SoftwareInfo SoftwareInfo  `json:"softwareInfo"`
}

func GetDeviceInfo() DeviceInfo {
	deviceInfo := DeviceInfo{
		UUID:         configure.GetUUID(),
		NetworkInfo:  GetListOfIPv4Interfaces(),
		SoftwareInfo: GetSoftwareInfo(),
	}

	return deviceInfo
}

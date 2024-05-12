package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/pkg/util/log"
)

type DeviceInfo struct {
	UUID         string        `json:"uuid"`
	NetworkInfo  []NetworkInfo `json:"networkInfo"`
	SoftwareInfo SoftwareInfo  `json:"softwareInfo"`
}

type ManagingDeviceDto struct {
	UUID string `json:"uuid"`
}

func GetDeviceInfo() DeviceInfo {
	deviceInfo := DeviceInfo{
		UUID:         configure.GetUUID(),
		NetworkInfo:  GetListOfIPv4Interfaces(),
		SoftwareInfo: GetSoftwareInfo(),
	}

	return deviceInfo
}

func ManageAccessAndDevices() (data string, err error) {
	_, _ = RegisterDevice()
	data, err = ManagingDevice()
	return data, err
}

func RegisterDevice() (data string, err error) {
	var url = GetApiHost() + "/devices"
	reqBody, err := json.Marshal(GetDeviceInfo())
	if err != nil {
		panic(err)
	}

	resp, err := httpPost(url, reqBody)
	if err != nil {
		err = fmt.Errorf("%w: %v", FailCreate, err)
		log.Warn("Device registration: %v", err)
		return "", err
	}

	log.Info("Device registration: %v -> SUCCESS", resp)
	return resp, err
}

func ManagingDevice() (data string, err error) {
	var url = GetApiHost() + "/me/manage_devices"
	var manageDeviceDto = ManagingDeviceDto{
		UUID: configure.GetUUID(),
	}
	reqBody, err := json.Marshal(manageDeviceDto)
	if err != nil {
		panic(err)
	}

	resp, err := httpPost(url, reqBody)
	if err != nil {
		err = fmt.Errorf("%w: %v", FailCreate, err)
		log.Warn("Managing device: %v", err)
		return "", err
	}

	log.Info("Managing device: %v -> SUCCESS", resp)
	return resp, err
}

func GetActivatedDevice() (data string, err error) {
	var url = GetApiHost() + "/devices/" + configure.GetUUID()
	resp, err := httpGet(url)
	if err != nil {
		err = fmt.Errorf("%w: %v", FailGet, err)
		log.Warn("GetActivatedDevice: %v", err)
		return "", err
	}

	return resp, err
}

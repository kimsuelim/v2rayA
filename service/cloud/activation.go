package cloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/v2rayA/v2rayA/common/httpClient"
	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/pkg/util/log"
)

type ManagingDeviceDto struct {
	UUID string `json:"uuid"`
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

func httpPost(url string, body []byte) (data string, err error) {
	client := httpClient.GetHttpClientAutomatically()
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+configure.GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// An error is returned if caused by client policy (such as CheckRedirect), or
	// failure to speak HTTP (such as a network connectivity problem).
	// A non-2xx status code doesn't cause an error.
	// ref: https://github.com/golang/go/issues/21846
	if resp.StatusCode >= 400 {
		return "", errors.New(string(b))
	}

	return string(b), nil
}

func httpGet(url string) (data string, err error) {
	client := httpClient.GetHttpClientAutomatically()
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	req.Header.Add("Authorization", "Bearer "+configure.GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", errors.New(string(b))
	}

	return string(b), nil
}

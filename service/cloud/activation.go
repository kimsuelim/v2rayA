package cloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/v2rayA/v2rayA/common/httpClient"
	"github.com/v2rayA/v2rayA/conf"
	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/pkg/util/log"
	"io"
)

var (
	FailCreate = fmt.Errorf("failed to create device")
	FailGet    = fmt.Errorf("failed to get device")
)

func ActivateDevice() (data string, err error) {
	log.Alert("Activating Device...")

	var url string
	if conf.IsDebug() {
		url = "http://host.docker.internal:8080/device"
	} else {
		url = "https://api.imcvpn.com/device"
	}

	reqBody, err := json.Marshal(GetDeviceInfo())
	if err != nil {
		panic(err)
	}

	resp, err := httpPost(url, reqBody)
	if err != nil {
		err = fmt.Errorf("%w: %v", FailCreate, err)
		log.Warn("Activating Device: %v", err)
		return "", err
	}

	log.Info("Activating Device: %v -> SUCCESS\n", resp)
	return
}

func GetActivatedDevice() (data string, err error) {
	var url string
	if conf.IsDebug() {
		url = "http://host.docker.internal:8080/device/" + configure.GetUUID()
	} else {
		url = "https://api.imcvpn.com/device" + configure.GetUUID()
	}

	resp, err := httpGet(url)
	if err != nil {
		err = fmt.Errorf("%w: %v", FailGet, err)
		log.Warn("GetActivatedDevice: %v", err)
		return "", err
	}

	log.Info("GetActivatedDevice: %v -> SUCCESS\n", resp)
	return
}

func httpPost(url string, body []byte) (data string, err error) {
	client := httpClient.GetHttpClientAutomatically()
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(body))
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
	resp, err := httpClient.GetHttpClientAutomatically().Get(url)
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

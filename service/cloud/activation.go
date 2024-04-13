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
	"net/http"
)

func ActivateDevice() (data string, err error) {
	var url = GetApiHost() + "/device"
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

	log.Info("Activating Device: %v -> SUCCESS", resp)
	return
}

func GetActivatedDevice() (data string, err error) {
	var url = GetApiHost() + "/device/" + configure.GetUUID()
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
	req.SetBasicAuth(conf.BasicAuthUsername, conf.BasicAuthPassword)

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
	req.SetBasicAuth(conf.BasicAuthUsername, conf.BasicAuthPassword)

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

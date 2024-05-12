package cloud

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/v2rayA/v2rayA/common/httpClient"
	"github.com/v2rayA/v2rayA/db/configure"
)

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

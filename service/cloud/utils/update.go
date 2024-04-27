package utils

import (
	"bytes"
	"fmt"
	"github.com/v2rayA/v2rayA/common"
	"github.com/v2rayA/v2rayA/conf"
	"net/http"
	"strings"
)

func CheckLatestUpdate() (foundNew bool, remoteVersion string, err error) {
	resp, err := http.Get("https://api.github.com/repos/kimsuelim/v2raya/releases/latest")
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(resp.Body)
	if err != nil && n > 0 {
		return
	}
	defer resp.Body.Close()
	s := buf.String()
	prefix := "\"name\":\"v"
	l := strings.Index(s, prefix)
	if l < 0 {
		return false, "", fmt.Errorf("failed to get latest version from github release")
	}
	s = s[l+len(prefix):]
	r := strings.Index(s, "\"")
	s = s[:r]
	ge, err := common.VersionGreaterEqual(conf.Version, s)
	return !ge, s, err
}

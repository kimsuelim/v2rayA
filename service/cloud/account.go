package cloud

import (
	"encoding/json"
	"fmt"
	"github.com/v2rayA/v2rayA/pkg/server/jwt"
	"github.com/v2rayA/v2rayA/pkg/util/log"
	"time"
)

func Login(username, password string) (token string, err error) {
	if !IsValidAccount(username, password) {
		return "", fmt.Errorf("wrong username or password")
	}
	dur := 30 * 24 * time.Hour
	return jwt.MakeJWT(map[string]string{
		"uname": username,
	}, &dur)
}

func IsValidAccount(username, password string) bool {
	var user = map[string]interface{}{
		"name":     username,
		"email":    username,
		"password": password,
	}

	var url = GetApiHost() + "/login"
	reqBody, err := json.Marshal(user)
	if err != nil {
		return false
	}

	resp, err := httpPost(url, reqBody)
	if err != nil {
		err = fmt.Errorf("%w: %v", FailCreate, err)
		log.Warn("IsValidAccount: %v", err)
		return false
	}

	log.Info("IsValidAccount: %v -> SUCCESS", resp)

	return true
}

package cloud

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/pkg/server/jwt"
	"github.com/v2rayA/v2rayA/pkg/util/log"
)

type LoginResponseDto struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiredAt"`
}

func Login(username, password string) (token string, err error) {
	valid, token := IsValidAccount(username, password)
	if !valid {
		return "", fmt.Errorf("wrong username or password")
	}

	configure.SetAccessToken(token)

	// Run in background lightweight thread
	go func() {
		resp, err := ManageAccessAndDevices()
		if err != nil {
			log.Error("Manage Access and Devices: %v -> FAIL", err)
			return
		}

		log.Info("Manage Access and Devices: %v -> SUCCESS", resp)
	}()

	// Run in synchronize Go routine
	err = SyncServerWithCloud()
	if err != nil {
		// rescue error softly.
		// Try to sync in background Go routine.
		log.Error("[SyncServers] updating server list: %v -> FAIL", err)
	} else {
		log.Info("[SyncServers] updating server list: -> SUCCESS")
	}

	dur := 30 * 24 * time.Hour
	return jwt.MakeJWT(map[string]string{
		"uname": username,
	}, &dur)
}

func IsValidAccount(username, password string) (bool, string) {
	var user = map[string]interface{}{
		"email":    username,
		"password": password,
	}

	var url = GetApiHost() + "/login"
	reqBody, err := json.Marshal(user)
	if err != nil {
		return false, ""
	}

	resp, err := httpPost(url, reqBody)
	if err != nil {
		err = fmt.Errorf("%w: %v", FailCreate, err)
		log.Warn("IsValidAccount: %v", err)
		return false, ""
	}

	loginDto := LoginResponseDto{}
	err = json.Unmarshal([]byte(resp), &loginDto)
	if err != nil {
		return false, ""
	}

	return true, loginDto.Token
}

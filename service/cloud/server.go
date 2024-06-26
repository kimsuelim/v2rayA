package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/v2rayA/v2rayA/core/serverObj"
	"github.com/v2rayA/v2rayA/db/configure"
	"github.com/v2rayA/v2rayA/pkg/util/log"
	"github.com/v2rayA/v2rayA/server/service"
)

// format description
// https://github.com/2dust/v2rayN/wiki/%E5%88%86%E4%BA%AB%E9%93%BE%E6%8E%A5%E6%A0%BC%E5%BC%8F%E8%AF%B4%E6%98%8E(ver-2)
type VmessConfig struct {
	Ps            string `json:"ps"`   // alias name
	Add           string `json:"add"`  // host address, IP or Domain
	Port          string `json:"port"` // port number
	ID            string `json:"id"`   // unique id
	Aid           string `json:"aid"`  // alter id
	Security      string `json:"scy"`  // security
	Net           string `json:"net"`  // transport protocol
	Type          string `json:"type"`
	Host          string `json:"host"`
	Path          string `json:"path"`
	TLS           string `json:"tls"`
	AllowInsecure bool   `json:"allowInsecure"`
	V             string `json:"v"`
	Protocol      string `json:"protocol"`
}

type ServerDto struct {
	Id        int         `json:"id"`
	Protocol  string      `json:"protocol"`
	Config    VmessConfig `json:"config"`
	ShareUrl  string      `json:"shareUrl"`
	Country   string      `json:"country"`
	State     string      `json:"state"`
	CreatedAt string      `json:"createdAt"`
	UpdatedAt string      `json:"updatedAt"`
}

func SyncServerWithCloud() (err error) {
	var localServers = configure.GetServers()
	var cloudServers []ServerDto

	resp, err := GetServerFromCloud()
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(resp), &cloudServers)
	if err != nil {
		return
	}

	for _, serverDto := range cloudServers {
		var obj serverObj.ServerObj
		obj, err = service.ResolveURL(serverDto.ShareUrl)
		if err != nil {
			continue
		}

		var isAdded = false
		for i, addedServer := range localServers {
			if (obj.GetHostname() == addedServer.ServerObj.GetHostname()) &&
				(obj.ProtoToShow() == addedServer.ServerObj.ProtoToShow()) &&
				(obj.GetPort() == addedServer.ServerObj.GetPort()) {
				//log.Info("SyncServerWithCloud: isAdded, index: %v", i)

				if !isOperational(serverDto.State) {
					serversIndexes := make([]int, 0, 1)
					serversIndexes = append(serversIndexes, i)
					err = configure.RemoveServers(serversIndexes)
					if err != nil {
						break
					}
					log.Info("SyncServerWithCloud: RemoveServers: %v, index: %v", serverDto.State, i)
				}

				isAdded = true
				break
			}
		}

		if isAdded == false {
			if isOperational(serverDto.State) {
				err = configure.AppendServers([]*configure.ServerRaw{{ServerObj: obj}})
				if err != nil {
					continue
				}
				log.Info("SyncServerWithCloud: AppendServers: %v", obj)
			}
		}
	}

	return
}

func GetServerFromCloud() (data string, err error) {
	var url = GetApiHost() + "/me/servers"
	resp, err := httpGet(url)
	if err != nil {
		err = fmt.Errorf("%w: %v", FailGet, err)
		return "", err
	}

	return resp, err
}

func isOperational(state string) bool {
	return state == "RUNNING"
}

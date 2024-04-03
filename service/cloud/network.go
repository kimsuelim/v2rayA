package cloud

import (
	"fmt"
	"net"
)

type NetworkInfo struct {
	Name         string `json:"name"`
	HardwareAddr string `json:"hardwareAddr"`
	IPAddr       string `json:"ipAddr"`
}

func GetListOfIPv4Interfaces() []NetworkInfo {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var as []NetworkInfo

	// get list of interfaces
	for _, i := range interfaces {
		// if mac address not exists, bypass it.
		if i.HardwareAddr == nil {
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			fmt.Printf("failed to get addresses of interface: %s, err: %w\n", i.Name, err)
			continue
		}

		// get IPv4 address per interface
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				// check if IPv4 or IPv6 is not nil
				//if ipnet.IP.To4() != nil || ipnet.IP.To16 != nil {

				// check if IPv4 is not nil
				if ipnet.IP.To4() != nil {
					// print available addresses
					//fmt.Printf("name: %s, ip: %s, mac: %s\n", i.Name, ipnet.IP, i.HardwareAddr)
					as = append(as, NetworkInfo{Name: i.Name, HardwareAddr: i.HardwareAddr.String(), IPAddr: ipnet.IP.String()})
				}
			}
		}
	}

	return as
}

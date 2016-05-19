package device42

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

var (
	updateIPAddress = "/ips/"
	nextIPAddress   = "/suggest_ip/"
)

// IPAddress :
type IPAddress struct {
	IP string `json:"ip"`
}

// GetNextIP :
func GetNextIP(subnet string) (string, error) {
	var ip IPAddress

	params := url.Values{}
	params.Add("subnet", subnet)

	paramData := params.Encode()

	data, err := makeRequest("POST", nextIPAddress, paramData)
	if err != nil {
		return "", err
	}

	readData, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(readData, &ip)
	if err != nil {
		return "", err
	}

	return ip.IP, nil
}

// ReserveIP :
func ReserveIP(ipaddress, device string) error {
	params := url.Values{}
	params.Add("ipaddress", ipaddress)
	params.Add("available", "no")
	params.Add("device", device)
	params.Add("label", device)

	paramData := params.Encode()

	_, err := makeRequest("POST", updateIPAddress, paramData)
	if err != nil {
		return err
	}
	return nil
}

// UpdateIP :
func UpdateIP(name, ip string) error {
	params := url.Values{}
	params.Add("name", name)
	params.Add("ip", ip)

	paramData := params.Encode()

	_, err := makeRequest("POST", updateIPAddress, paramData)
	if err != nil {
		return err
	}
	return nil
}

// MakeIPAvailable :
func MakeIPAvailable(ipaddress string) error {
	params := url.Values{}
	params.Add("ipaddress", ipaddress)
	params.Add("clear_all", "yes")

	paramData := params.Encode()

	_, err := makeRequest("POST", updateIPAddress, paramData)
	if err != nil {
		return err
	}
	return nil
}

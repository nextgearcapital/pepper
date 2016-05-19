package device42

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/nextgearcapital/pepper/pkg/log"
)

// Device :
type Device struct {
	ID string `json:"id"`
}

var (
	getDeviceID  = "/devices/name/"
	createDevice = "/device/"
	deleteDevice = "/devices/"
)

func makeRequest(method, endpoint, data string) (io.ReadCloser, error) {
	// This is only here temporarily until the binary is on the salt master with the proper certs
	// Ignore SSL cert
	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	noTLS := &http.Client{Transport: transportConfig}

	var request *http.Request
	var err error

	if method == "GET" {
		request, err = http.NewRequest(method, BaseURL+endpoint, nil)
		if err != nil {
			return nil, err
		}
	} else {
		request, err = http.NewRequest(method, BaseURL+endpoint, strings.NewReader(data))
		if err != nil {
			return nil, err
		}
	}
	// This is the encoding that d42 requires because their JSON API sucks
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(Username, Password)
	response, err := noTLS.Do(request)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

// CleanDeviceAndIP :
func CleanDeviceAndIP(host, ipAddress string) error {
	if err := MakeIPAvailable(ipAddress); err != nil {
		return err
	}
	if err := DeleteDevice(host); err != nil {
		return err
	}
	return nil
}

// CreateDevice :
func CreateDevice(host string) error {
	params := url.Values{}
	params.Add("name", host)
	params.Add("service_level", ServiceLevel)
	params.Add("type", "virtual")
	params.Add("virtual_subtype", "vmware")

	paramData := params.Encode()

	_, err := makeRequest("POST", createDevice, paramData)
	if err != nil {
		return err
	}
	return nil
}

// This only exists because you can't delete things by name
// so we need to grab the ID by doing this
func getDevice(host string) (string, error) {
	var d Device

	params := url.Values{}
	params.Add("name", host)

	paramData := params.Encode()

	data, err := makeRequest("GET", getDeviceID, paramData)
	if err != nil {
		return "", err
	}

	readData, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(readData, &d)
	if err != nil {
		return "", err
	}

	return d.ID, nil
}

// DeleteDevice :
func DeleteDevice(host string) error {
	id, err := getDevice(host)
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Add("id", string(id))

	paramData := params.Encode()

	_, err = makeRequest("DELETE", deleteDevice, paramData)
	if err != nil {
		return err
	}
	log.Err(paramData)
	return nil
}

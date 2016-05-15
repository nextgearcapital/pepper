package device42

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Device : Satisfy golint
type Device struct {
	ID int `json:"id"`
}

var (
	getDeviceID  = BaseURL + "/devices/name/"
	createDevice = BaseURL + "/device/"
	deleteDevice = BaseURL + "/devices/"
)

func makeRequest(verb, endpoint, data string) (io.ReadCloser, error) {
	// This is only here temporarily until the binary is on the salt master with the proper certs
	// Ignore SSL cert
	transportConfig := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	noTLS := &http.Client{Transport: transportConfig}

	var request *http.Request
	var err error

	if verb != "GET" {
		request, err = http.NewRequest(verb, endpoint, strings.NewReader(data))
		if err != nil {
			return nil, err
		}
	} else if verb == "GET" {
		request, err = http.NewRequest(verb, endpoint, nil)
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

// CreateDevice : Satisfy golint
func CreateDevice(host string, servicelevel string) error {
	params := url.Values{}
	params.Add("name", host)
	params.Add("service_level", servicelevel)
	params.Add("type", "virtual")
	params.Add("virtual_subtype", "vmware")

	paramData := params.Encode()

	_, err := makeRequest(paramData, createDevice, "POST")
	if err != nil {
		return err
	}
	return nil
}

// GetDevice : Satisfy golint
func GetDevice(host string) (int, error) {
	var d Device

	params := url.Values{}
	params.Add("name", host)

	paramData := params.Encode()

	data, err := makeRequest(paramData, getDeviceID, "GET")
	if err != nil {
		return -1, err
	}

	readData, err := ioutil.ReadAll(data)
	if err != nil {
		return -1, err
	}

	err = json.Unmarshal(readData, &d)
	if err != nil {
		return -1, err
	}

	return d.ID, nil
}

// DeleteDevice : Satisfy golint
func DeleteDevice(host string) error {
	id, err := GetDevice(host)
	if err != nil {
		return err
	}

	realID := strconv.Itoa(id)

	params := url.Values{}
	params.Add("id", realID)

	paramData := params.Encode()

	_, err = makeRequest(paramData, deleteDevice, "DELETE")
	if err != nil {
		return err
	}
	return nil
}

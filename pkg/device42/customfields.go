package device42

import (
	"net/url"
)

var (
	updateCustomFields = "/device/custom_field/"
)

// UpdateCustomFields :
func UpdateCustomFields(host, key, value string) error {
	params := url.Values{}
	params.Add("name", host)
	params.Add("key", key)
	params.Add("value", value)

	paramData := params.Encode()

	_, err := makeRequest("PUT", updateCustomFields, paramData)
	if err != nil {
		return err
	}

	return nil
}

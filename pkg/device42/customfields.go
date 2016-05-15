package device42

import (
	"net/url"
)

var (
	updateCustomFields = BaseURL + "/device/custom_field/"
)

// UpdateCustomFields : Satisfy golint
func UpdateCustomFields(host, key, value string) error {
	params := url.Values{}
	params.Add("name", host)
	params.Add("key", key)
	params.Add("value", value)

	paramData := params.Encode()

	_, err := makeRequest(paramData, updateCustomFields, "PUT")
	if err != nil {
		return err
	}

	return nil
}

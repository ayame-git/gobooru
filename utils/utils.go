package utils

import (
	"bytes"
	"encoding/json"
	"net/url"
)

func ToBuffer(form url.Values) (*bytes.Buffer, error) {
	data, err := json.Marshal(form)
	return bytes.NewBuffer(data), err
}

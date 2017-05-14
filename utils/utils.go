package utils

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

//reason to use url.Values instead of map[string]string
//is it's easier to work with tags
//...i think
func toQuery(v url.Values) map[string]string {
	q := make(map[string]string)
	for key, val := range v {
		str := strings.Join(val, " ")
		q[key] = str
	}
	return q
}

func ToBuffer(v url.Values) (*bytes.Buffer, error) {
	q := toQuery(v)
	data, err := json.Marshal(q)
	return bytes.NewBuffer(data), err
}

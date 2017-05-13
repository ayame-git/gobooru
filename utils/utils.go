package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"os"
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

func ToForm(v url.Values, field string) (*bytes.Buffer, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	defer w.Close()

	q := toQuery(v)
	for key, val := range q {
		var fieldname string
		if key == "login" || key == "password_hash" {
			fieldname = key
		} else {
			fieldname = fieldName(field, key)
		}

		if key == "file" {
			f, err := os.Open(val)
			if err != nil {
				return nil, err
			}

			defer f.Close()

			fw, err := w.CreateFormFile(fieldname, val)
			if err != nil {
				return nil, err
			}
			if _, err = io.Copy(fw, f); err != nil {
				return nil, err
			}
		} else {
			err := w.WriteField(fieldname, val)
			if err != nil {
				return nil, err
			}
		}
	}
	return &b, nil
}

func fieldName(field, name string) string {
	return field + "[" + name + "]"
}

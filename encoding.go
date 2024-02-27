// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import (
	"encoding/base64"
	"encoding/json"
	"html/template"

	"sigs.k8s.io/yaml"
)

// EncodeBase64 returns the base64 encoding of s.
func EncodeBase64(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString([]byte(ss)), nil
}

// DecodeBase64 returns the base64 decoding of s.
func DecodeBase64(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	d, err := base64.StdEncoding.DecodeString(ss)
	return string(d), err
}

// EncodeJSON returns the JSON encoding of v.
func EncodeJSON(v interface{}) (template.HTML, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.HTML(b), nil
}

// DecodeJSON returns the JSON decoding of s.
func DecodeJSON(s interface{}) (map[string]interface{}, error) {
	v := map[string]interface{}{}

	data, err := toStringE(s)
	if err != nil {
		return v, err
	}
	err = json.Unmarshal([]byte(data), &v)
	return v, err
}

// EncodeYAML returns the YAML encoding of v.
func EncodeYAML(v interface{}) (template.HTML, error) {
	b, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.HTML(b), nil
}

// DecodeYAML returns the YAML decoding of s.
func DecodeYAML(s interface{}) (map[string]interface{}, error) {
	v := map[string]interface{}{}

	data, err := toStringE(s)
	if err != nil {
		return v, err
	}
	err = yaml.Unmarshal([]byte(data), &v)
	return v, err
}

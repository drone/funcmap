// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import "testing"

func TestEncodeJson(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"name": "jane"},
			text: `{{ jsonify . }}`,
			want: `{"name":"jane"}`,
		},
	}
	for _, test := range tests {
		out, err := execute(test.text, test.data)
		if err != nil {
			t.Error(err)
		} else if out != test.want {
			t.Errorf("Want template result %s, got %s", test.want, out)
		}
	}
}

func TestDecodeJson(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"body": `{"name":"jane"}`},
			text: `{{ (jsonDecode .body).name }}`,
			want: `jane`,
		},
	}
	for _, test := range tests {
		out, err := execute(test.text, test.data)
		if err != nil {
			t.Error(err)
		} else if out != test.want {
			t.Errorf("Want template result %s, got %s", test.want, out)
		}
	}
}

func TestEncodeYaml(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"name": "jane"},
			text: `{{ yamlEncode . }}`,
			want: "name: jane\n",
		},
	}
	for _, test := range tests {
		out, err := execute(test.text, test.data)
		if err != nil {
			t.Error(err)
		} else if out != test.want {
			t.Errorf("Want template result %s, got %s", test.want, out)
		}
	}
}

func TestDecodeYaml(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"body": `name: jane`},
			text: `{{ (yamlDecode .body).name }}`,
			want: `jane`,
		},
	}
	for _, test := range tests {
		out, err := execute(test.text, test.data)
		if err != nil {
			t.Error(err)
		} else if out != test.want {
			t.Errorf("Want template result %s, got %s", test.want, out)
		}
	}
}

func TestEncodeBase64(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"name": "jane"},
			text: `{{ base64Encode .name }}`,
			want: "amFuZQ==",
		},
	}
	for _, test := range tests {
		out, err := execute(test.text, test.data)
		if err != nil {
			t.Error(err)
		} else if out != test.want {
			t.Errorf("Want template result %s, got %s", test.want, out)
		}
	}
}

func TestDecodeBase64(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"name": "amFuZQ=="},
			text: `{{ base64Decode .name }}`,
			want: "jane",
		},
	}
	for _, test := range tests {
		out, err := execute(test.text, test.data)
		if err != nil {
			t.Error(err)
		} else if out != test.want {
			t.Errorf("Want template result %s, got %s", test.want, out)
		}
	}
}

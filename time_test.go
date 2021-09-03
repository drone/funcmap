// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import "testing"

func TestTime(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"Timestamp": "2021-09-03 07:13:02 -0700 MST"},
			text: `{{ time .Timestamp }}`,
			want: "2021-09-03 07:13:02 -0700 MST",
		},
		{
			data: map[string]interface{}{"Timestamp": 1630678382},
			text: `{{ time .Timestamp "UTC" }}`,
			want: "2021-09-03 14:13:02 +0000 UTC",
		},
		{
			data: map[string]interface{}{"Timestamp": 1630678382},
			text: `{{ time .Timestamp "MST" }}`,
			want: "2021-09-03 07:13:02 -0700 MST",
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

func TestTimeFormat(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"Timestamp": "2021-09-03 07:13:02 -0700 MST"},
			text: `{{ dateFormat "2006-01-02" .Timestamp }}`,
			want: "2021-09-03",
		},
		{
			data: map[string]interface{}{"Timestamp": 1630678382},
			text: `{{ dateFormat "2006-01-02" .Timestamp }}`,
			want: "2021-09-03",
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

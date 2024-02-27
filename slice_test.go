// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import "testing"

func TestSlices(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{}
	for _, test := range tests {
		out, err := execute(test.text, test.data)
		if err != nil {
			t.Error(err)
		} else if out != test.want {
			t.Errorf("Want template result %s, got %s", test.want, out)
		}
	}
}

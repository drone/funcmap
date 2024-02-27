// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import "testing"

func TestCrypto(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		{
			data: map[string]interface{}{"input": "hello world"},
			text: `{{ md5 .input }}`,
			want: `5eb63bbbe01eeed093cb22bb8f5acdc3`,
		},
		{
			data: map[string]interface{}{"input": "hello world"},
			text: `{{ sha1 .input }}`,
			want: `2aae6c35c94fcfb415dbe95f408b9ce91ee846ed`,
		},
		{
			data: map[string]interface{}{"input": "hello world"},
			text: `{{ sha256 .input }}`,
			want: `b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9`,
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

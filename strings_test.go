// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import "testing"

func TestStrings(t *testing.T) {
	tests := []struct {
		data map[string]interface{}
		text string
		want string
	}{
		//
		// append / prepend
		//
		{
			data: nil,
			text: `{{ append "hello" "world" }}`,
			want: "helloworld",
		},
		{
			data: nil,
			text: `{{ prepend "hello" "world" }}`,
			want: "worldhello",
		},
		//
		// chomp
		//
		{
			data: map[string]interface{}{"input": "hello world\n"},
			text: `{{ chomp .input }}`,
			want: "hello world",
		},
		//
		// contains
		//
		{
			data: nil,
			text: `{{ contains "foobar" "foo" }}`,
			want: "true",
		},
		{
			data: nil,
			text: `{{ contains "foobar" "baz" }}`,
			want: "false",
		},
		//
		// findRE
		//
		{
			data: nil,
			text: `{{ findRE "foo" "foobar" }}`,
			want: "[foo]",
		},
		//
		// hasPrefix
		//
		{
			data: nil,
			text: `{{ hasPrefix "foobar" "foo" }}`,
			want: "true",
		},
		{
			data: nil,
			text: `{{ hasPrefix "foobar" "bar" }}`,
			want: "false",
		},
		//
		// hasSuffix
		//
		{
			data: nil,
			text: `{{ hasSuffix "foobar" "bar" }}`,
			want: "true",
		},
		{
			data: nil,
			text: `{{ hasSuffix "foobar" "foo" }}`,
			want: "false",
		},
		//
		// padLeft
		//
		{
			data: nil,
			text: `{{ padLeft "hello" " " 5 }}`,
			want: "     hello",
		},
		//
		// padRight
		//
		{
			data: nil,
			text: `{{ padRight "hello" " " 5 }}`,
			want: "hello     ",
		},
		//
		// repeat
		//
		{
			data: nil,
			text: `{{ repeat " " 5 }}`,
			want: "     ",
		},
		//
		// replace
		//
		{
			data: nil,
			text: `{{ replace "foobazbar" "baz" " " }}`,
			want: "foo bar",
		},
		{
			data: nil,
			text: `{{ replaceRE "^https?://([^/]+).*" "$1" "https://drone.io" }}`,
			want: "drone.io",
		},
		//
		// lower / title / upper
		//
		{
			data: nil,
			text: `{{ lower "HELLO" }}`,
			want: "hello",
		},
		{
			data: nil,
			text: `{{ title "hello" }}`,
			want: "HELLO",
		},
		{
			data: nil,
			text: `{{ upper "hello" }}`,
			want: "HELLO",
		},
		//
		// split
		//
		{
			data: nil,
			text: `{{ split "foo,bar,baz" "," }}`,
			want: "[foo bar baz]",
		},
		{
			data: nil,
			text: `{{ splitn "foo,bar,baz" "," 2 }}`,
			want: "[foo bar,baz]",
		},
		//
		// trim
		//
		{
			data: nil,
			text: `{{ trim " foo " }}`,
			want: "foo",
		},
		{
			data: nil,
			text: `{{ trimLeft " foo " " " }}`,
			want: "foo ",
		},
		{
			data: nil,
			text: `{{ trimRight " foo " " " }}`,
			want: " foo",
		},
		{
			data: nil,
			text: `{{ trimPrefix " foo " " " }}`,
			want: "foo ",
		},
		{
			data: nil,
			text: `{{ trimSuffix " foo " " " }}`,
			want: " foo",
		},
		//
		// urlize
		//
		{
			data: nil,
			text: `{{ urlize "foo bar baz" }}`,
			want: "foo-bar-baz",
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

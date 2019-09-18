// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import (
	"bytes"
	"testing"
	"text/template"
)

func TestSomething(t *testing.T) {
	buf := new(bytes.Buffer)
	tmpl, _ := template.New("_").Funcs(Funcs).Parse(testT)
	err := tmpl.Execute(buf, map[string]interface{}{
		"Num": float64(32.5),
	})
	if err != nil {
		t.Error(err)
	}
	println(buf.String())
}

var testT = `{{ printf "hello %s %s world" (upper "foo" ) (lower .Num ) }}`

// https://github.com/leekchan/gtf
// https://github.com/Masterminds/sprig

// # HUGO

// # MISC
// default
// seq       (collections?) slices
// cond

// # URL
// querify  (collections?) slices

// CUSTOM
// * uppercasefirst
// * since
// * base, dir, ext, clean, abs

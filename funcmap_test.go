// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import (
	"bytes"
	"text/template"
)

// helper function parses and applies the template and returns the
// results in string format.
func execute(text string, data map[string]interface{}) (string, error) {
	var buf bytes.Buffer
	t, err := template.New("_").Funcs(Funcs).Parse(text)
	if err != nil {
		return buf.String(), err
	}
	err = t.Execute(&buf, data)
	return buf.String(), err
}

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

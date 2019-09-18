// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import (
	"html"
)

// HTMLEscape returns a copy of s with reserved HTML
// characters escaped.
func HTMLEscape(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return html.EscapeString(ss), nil
}

// HTMLUnescape returns a copy of s with HTML escape
// sequences converted to plain text.
func HTMLUnescape(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return html.UnescapeString(ss), nil
}

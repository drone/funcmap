// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import "time"

// Time converts the textual representation of the datetime
// string into a time.Time interface.
func Time(v interface{}) (interface{}, error) {
	t, err := toTimeE(v)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// TimeFormat converts the textual representation of the
// datetime string into the other form or returns it of
// the time.Time value. These are formatted with the layout
// string
func TimeFormat(layout string, v interface{}) (string, error) {
	t, err := toTimeE(v)
	if err != nil {
		return "", err
	}
	return t.Format(layout), nil
}

// Now returns the current local time.
func Now() time.Time {
	return time.Now()
}

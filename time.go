// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import "time"

// Time converts the textual representation of the datetime
// string into a time.Time interface.
func Time(v interface{}, args ...interface{}) (interface{}, error) {
	t, err := toTimeE(v)
	if err != nil {
		return nil, err
	}

	if len(args) > 0 {
		locStr, err := toStringE(args[0])
		if err != nil {
			return nil, err
		}
		loc, err := time.LoadLocation(locStr)
		if err != nil {
			return nil, err
		}
		return t.In(loc), nil
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

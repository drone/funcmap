// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 returns the sha checksum of s.
func MD5(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	hash := md5.Sum([]byte(ss))
	return hex.EncodeToString(hash[:]), nil
}

// SHA1 returns the sha checksum of s.
func SHA1(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	hash := sha1.Sum([]byte(ss))
	return hex.EncodeToString(hash[:]), nil
}

// SHA256 returns the sha checksum of s.
func SHA256(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256([]byte(ss))
	return hex.EncodeToString(hash[:]), nil
}

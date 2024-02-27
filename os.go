// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import (
	"errors"
	"io/ioutil"
	"os"
)

// Getenv retrieves the value of the environment variable
// named by the key.
func Getenv(key interface{}) (string, error) {
	sk, err := toStringE(key)
	if err != nil {
		return "", err
	}
	return os.Getenv(sk), nil
}

// ReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadDir(dirname interface{}) ([]os.FileInfo, error) {
	sd, err := toStringE(dirname)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadDir(sd)
}

// ReadFile reads the directory named by dirname and returns
// a list of directory entries sorted by filename.
func ReadFile(filename interface{}) ([]byte, error) {
	sf, err := toStringE(filename)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(sf)
}

// FileExists returns true if the file exists.
func FileExists(name interface{}) (bool, error) {
	sn, err := toStringE(name)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(sn)
	return err == nil, nil
}

// Stat returns the os.FileInfo structure describing file.
func Stat(i interface{}) (os.FileInfo, error) {
	path, err := toStringE(i)
	if err != nil {
		return nil, err
	}

	if path == "" {
		return nil, errors.New("fileStat needs a path to a file")
	}

	r, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import (
	"net/url"
	"strings"
	"sigs.k8s.io/yaml"
)

// TODO (bradrydzewski) add Join function
// TODO (bradrydzewski) add Substr function
// TODO (bradrydzewski) add ReplaceRE function
// TODO (bradrydzewski) add FindRE function

// Append returns a slice of the string s, with append appended.
func Append(s, append interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	as, err := toStringE(append)
	if err != nil {
		return "", err
	}
	return ss + as, nil
}

// Chomp returns a copy of s with all trailing newline
// characters removed.
func Chomp(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(ss, "\r\n"), nil
}

// Contains reports whether substr is within s.
func Contains(s, substr interface{}) (bool, error) {
	ss, err := toStringE(s)
	if err != nil {
		return false, err
	}
	su, err := toStringE(substr)
	if err != nil {
		return false, err
	}
	return strings.Contains(ss, su), nil
}

// ContainsAny reports whether any Unicode code points in
// chars are within s.
func ContainsAny(s, chars interface{}) (bool, error) {
	ss, err := toStringE(s)
	if err != nil {
		return false, err
	}
	sc, err := toStringE(chars)
	if err != nil {
		return false, err
	}
	return strings.ContainsAny(ss, sc), nil
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s interface{}, prefix string) (bool, error) {
	ss, err := toStringE(s)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(ss, prefix), nil
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s interface{}, suffix string) (bool, error) {
	ss, err := toStringE(s)
	if err != nil {
		return false, err
	}
	return strings.HasSuffix(ss, suffix), nil
}

// Indent returns a copy of the string s, with all lines prefixed with n spaces
func Indent(s interface{}, n int) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	pad := strings.Repeat(" ", n)
	return pad + strings.Replace(ss, "\n", "\n"+pad, -1), nil
}

// NIndent returns a copy of the string s, with all lines prefixed with n
// spaces, and a newline at the start. This is useful when interpolating into
// yaml, where you want each line indented by the exact same amount, but 
func NIndent(s interface{}, n int) (string, error) {
	indented, err := Indent(s, n)
	if err != nil {
		return "", err
	}
	return "\n" + indented, nil
}

// Prepend returns a slice of the string s, prepended with prepend.
func Prepend(s, prepend interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	ps, err := toStringE(prepend)
	if err != nil {
		return "", err
	}
	return ps + ss, nil
}

// PadLeft returns a slice of the string s, prefixed with
// n copies of padding.
func PadLeft(s, padding interface{}, n int) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	sp, err := toStringE(padding)
	if err != nil {
		return "", err
	}
	for i := 0; i < n; i++ {
		ss = sp + ss
	}
	return ss, err
}

// PadRight returns a slice of the string s, suffixed with
// n instances of the padding string.
func PadRight(s, padding interface{}, n int) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	sp, err := toStringE(padding)
	if err != nil {
		return "", err
	}
	for i := 0; i < n; i++ {
		ss = ss + sp
	}
	return ss, err
}

// Replace returns a copy of the string s with instances of
// old replaced by new.
func Replace(s, old, new interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	so, err := toStringE(old)
	if err != nil {
		return "", err
	}
	no, err := toStringE(new)
	if err != nil {
		return "", err
	}
	return strings.Replace(ss, so, no, -1), nil
}

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
func Split(s interface{}, sep string) ([]string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return nil, err
	}
	return strings.Split(ss, sep), nil
}

// SplitN slices s into substrings separated by sep and returns
// a slice of the substrings between those separators.
func SplitN(s interface{}, sep string, n int) ([]string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return nil, err
	}
	return strings.SplitN(ss, sep, n), nil
}

// ToLower returns a copy of the input s with all Unicode letters
// mapped to their lower case.
func ToLower(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.ToLower(ss), nil
}

// ToTitle returns a copy of the input s with all Unicode letters
// mapped to their title case.
func ToTitle(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(ss), nil
}

// ToUpper returns a copy of the input s with all Unicode letters
// mapped to their upper case.
func ToUpper(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(ss), nil
}

// TrimLeft returns a slice of the string s with all leading Unicode
// code points contained in cutset removed.
func TrimLeft(s interface{}, custet string) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.TrimLeft(ss, custet), nil
}

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
func TrimRight(s interface{}, custet string) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(ss, custet), nil
}

// TrimPrefix returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
func TrimPrefix(s interface{}, prefix string) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(ss, prefix), nil
}

// TrimSuffix returns s without the provided trailing suffix
// string. If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s interface{}, suffix string) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(ss, suffix), nil
}

// Trim returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func Trim(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(ss), nil
}

// Urlize returns a copy of the input s that is sanities so
// it can be safely placed inside a URL path segment, with
// spaces converted to hypens.
func Urlize(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	ss = strings.TrimSpace(ss)
	ss = strings.Replace(ss, " ", "-", -1)
	return url.QueryEscape(ss), nil
}

// toYAML takes an interface, marshals it to yaml, and returns a string. It will
// always return a string, even on marshal error (empty string).
//
// This is designed to be called from a template.
func toYAML(s interface{}) (string, error) {
	data, err := yaml.Marshal(s)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(data), "\n")
}

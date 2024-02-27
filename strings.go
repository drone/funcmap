// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import (
	"net/url"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// String functions from Hugo
// https://github.com/gohugoio/hugo/blob/7f47b99ea9a7ba7759c4f9424fedd4591e6da497/tpl/strings/strings.go

// TODO (bradrydzewski) add Join function
// TODO (bradrydzewski) add Substr function

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

// FindRE returns a list of strings that match the regular expression. By default all matches
// will be included. The number of matches can be limited with an optional third parameter.
func FindRE(expr string, content interface{}, limit ...interface{}) ([]string, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}

	conv, err := toStringE(content)
	if err != nil {
		return nil, err
	}

	if len(limit) == 0 {
		return re.FindAllString(conv, -1), nil
	}

	lim, err := toIntE(limit[0])
	if err != nil {
		return nil, err
	}

	return re.FindAllString(conv, lim), nil
}

// FirstUpper returns a string with the first character as upper case.
func FirstUpper(s interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}

	return firstUpper(ss), nil
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix interface{}) (bool, error) {
	s1, err := toStringE(s)
	if err != nil {
		return false, err
	}
	s2, err := toStringE(prefix)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(s1, s2), nil
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix interface{}) (bool, error) {
	s1, err := toStringE(s)
	if err != nil {
		return false, err
	}
	s2, err := toStringE(suffix)
	if err != nil {
		return false, err
	}
	return strings.HasSuffix(s1, s2), nil
}

// PadLeft returns a slice of the string s, prefixed with
// n copies of padding.
func PadLeft(s, padding, n interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	sp, err := toStringE(padding)
	if err != nil {
		return "", err
	}
	sn, err := toIntE(n)
	if err != nil {
		return "", err
	}
	for i := 0; i < sn; i++ {
		ss = sp + ss
	}
	return ss, err
}

// PadRight returns a slice of the string s, suffixed with
// n instances of the padding string.
func PadRight(s, padding, n interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	sp, err := toStringE(padding)
	if err != nil {
		return "", err
	}
	sn, err := toIntE(n)
	if err != nil {
		return "", err
	}
	for i := 0; i < sn; i++ {
		ss = ss + sp
	}
	return ss, err
}

// Repeat returns a new string consisting of count copies
// of the string s.
func Repeat(s, count interface{}) (string, error) {
	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}
	sn, err := toIntE(count)
	if err != nil {
		return "", err
	}
	return strings.Repeat(ss, sn), nil
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

// ReplaceRE returns a copy of s, replacing all matches of the regular
// expression pattern with the replacement text repl.
func ReplaceRE(pattern, repl, s interface{}) (string, error) {
	sp, err := toStringE(pattern)
	if err != nil {
		return "", err
	}

	sr, err := toStringE(repl)
	if err != nil {
		return "", err
	}

	ss, err := toStringE(s)
	if err != nil {
		return "", err
	}

	re, err := regexp.Compile(sp)
	if err != nil {
		return "", err
	}

	return re.ReplaceAllString(ss, sr), nil
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
func SplitN(s, sep, n interface{}) ([]string, error) {
	s1, err := toStringE(s)
	if err != nil {
		return nil, err
	}
	s2, err := toStringE(sep)
	if err != nil {
		return nil, err
	}
	no, err := toIntE(n)
	if err != nil {
		return nil, err
	}
	return strings.SplitN(s1, s2, no), nil
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
func TrimLeft(s, cutset interface{}) (string, error) {
	s1, err := toStringE(s)
	if err != nil {
		return "", err
	}
	s2, err := toStringE(cutset)
	if err != nil {
		return "", err
	}
	return strings.TrimLeft(s1, s2), nil
}

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
func TrimRight(s, cutset interface{}) (string, error) {
	s1, err := toStringE(s)
	if err != nil {
		return "", err
	}
	s2, err := toStringE(cutset)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(s1, s2), nil
}

// TrimPrefix returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
func TrimPrefix(s, prefix interface{}) (string, error) {
	s1, err := toStringE(s)
	if err != nil {
		return "", err
	}
	s2, err := toStringE(prefix)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(s1, s2), nil
}

// TrimSuffix returns s without the provided trailing suffix
// string. If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix interface{}) (string, error) {
	s1, err := toStringE(s)
	if err != nil {
		return "", err
	}
	s2, err := toStringE(suffix)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(s1, s2), nil
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

//
// helpers
//

func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

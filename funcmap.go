// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// license that can be found in the LICENSE file.

package funcmap

import "text/template"

// Funcs is a map of custom template functions.
var Funcs = map[string]interface{}{
	"append":       Append,
	"base64Decode": DecodeBase64,
	"base64Encode": EncodeBase64,
	"chomp":        Chomp,
	"contains":     Contains,
	"containsAny":  ContainsAny,
	"dateFormat":   TimeFormat,
	"env":          Getenv,
	"fileExists":   FileExists,
	"htmlEscape":   HTMLEscape,
	"htmlUnescape": HTMLUnescape,
	"hasPrefix":    HasPrefix,
	"hasSuffix":    HasSuffix,
	"indent":       Indent,
	"jsonify":      EncodeJSON,
	"lower":        ToLower,
	"md5":          MD5,
	"nindent":      NIndent,
	"now":          Now,
	"padLeft":      PadLeft,
	"padRight":     PadRight,
	"prepend":      Prepend,
	"readDir":      ReadDir,
	"readFile":     ReadFile,
	"replace":      Replace,
	"sha1":         SHA1,
	"sha256":       SHA256,
	"split":        Split,
	"splitn":       SplitN,
	"title":        ToTitle,
	"time":         Time,
	"trimLeft":     TrimLeft,
	"trimRight":    TrimRight,
	"trimPrefix":   TrimPrefix,
	"trimSuffix":   TrimSuffix,
	"trim":         Trim,
	"upper":        ToUpper,
	"urlize":       Urlize,
	"toYaml":       toYAML,
	"fromYaml":     fromYAML,
}

// SafeFuncs is a map of custom template functions. Functions
// that expose the environment and filesystem are excluded.
var SafeFuncs = map[string]interface{}{
	"append":       Append,
	"base64Decode": DecodeBase64,
	"base64Encode": EncodeBase64,
	"chomp":        Chomp,
	"contains":     Contains,
	"containsAny":  ContainsAny,
	"dateFormat":   TimeFormat,
	"htmlEscape":   HTMLEscape,
	"htmlUnescape": HTMLUnescape,
	"hasPrefix":    HasPrefix,
	"hasSuffix":    HasSuffix,
	"indent":       Indent,
	"jsonify":      EncodeJSON,
	"lower":        ToLower,
	"md5":          MD5,
	"nindent":      NIndent,
	"now":          Now,
	"padLeft":      PadLeft,
	"padRight":     PadRight,
	"prepend":      Prepend,
	"replace":      Replace,
	"sha1":         SHA1,
	"sha256":       SHA256,
	"split":        Split,
	"splitn":       SplitN,
	"title":        ToTitle,
	"time":         Time,
	"trimLeft":     TrimLeft,
	"trimRight":    TrimRight,
	"trimPrefix":   TrimPrefix,
	"trimSuffix":   TrimSuffix,
	"trim":         Trim,
	"upper":        ToUpper,
	"urlize":       Urlize,
	"toYaml":       toYAML,
	"fromYaml":     fromYAML,
}

// Combine combines function maps.
func Combine(funcmaps ...map[string]interface{}) template.FuncMap {
	result := map[string]interface{}{}
	for _, funcmap := range funcmaps {
		for k, v := range funcmap {
			result[k] = v
		}
	}
	return result
}

// after
// apply
// cond
// default
// delimit
// dict
// echoParam
// findRE
// first
// float
// in
// int
// intersect
// isset
// last
// pluralize
// querify
// range
// replaceRE
// seq
// shuffle
// slice
// slicestr
// sort
// string
// substr
// union
// uniq
// where

//
// ------------------------------------------
//

// WONT IMPLEMENT
// safeCSS
// safeHTML
// safeHTMLAttr
// safeJS
// safeURL
// plainify
// markdownify
// emojify
// countrunes
// countwords
// highlight
// singularize
// truncate
// humanize
// urls.Parse

// SPECIFIC TO HUGO
// ref
// relLangURL
// relURL
// relref
// i18n
// partialCached
// lang.NumFmt
// imageConfig
// render

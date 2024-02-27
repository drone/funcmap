This package aims to provide an independent implementation of the Hugo template functions.
https://gohugo.io/functions/ ([view source](https://github.com/gohugoio/hugo/tree/7f47b99ea9a7ba7759c4f9424fedd4591e6da497/tpl))

# Usage

Import the package:

```
import "github.com/drone/funcmap"
```

Provide the Funcmap when compiling your templates:

```
t, err := template.New("_").Funcs(funcmap.Funcs).Parse(text)
```

# Functions

* `append`
* `base64Decode`
* `base64Encode`
* `chomp`
* `contains`
* `containsAny`
* `dateFormat`
* `env`
* `findRE`
* `fileExists`
* `htmlEscape`
* `htmlUnescape`
* `hasPrefix`
* `hasSuffix`
* `jsonify`
* `jsonEncode`
* `jsonDecode`
* `lower`
* `md5`
* `now`
* `padLeft`
* `padRight`
* `prepend`
* `firstUpper`
* `readDir`
* `readFile`
* `replace`
* `replaceRE`
* `repeat`
* `sha1`
* `sha256`
* `slicestr`
* `split`
* `splitn`
* `title`
* `time`
* `trimLeft`
* `trimRight`
* `trimPrefix`
* `trimSuffix`
* `trim`
* `upper`
* `urlize`
* `yamlEncode`
* `yamlDecode`

Todo:

* `after`
* `apply`
* `cond`
* `default`
* `delimit`
* `dict`
* `echoParam`
* `first`
* `float`
* `in`
* `int`
* `intersect`
* `isset`
* `last`
* `pluralize`
* `querify`
* `range`
* `seq`
* `shuffle`
* `slice`
* `sort`
* `string`
* `substr`
* `union`
* `uniq`
* `where`

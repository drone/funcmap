This package aims to provide an independent implementation of the Hugo template functions.
https://gohugo.io/functions/

__Available Functions:__

* `append`
* `base64Decode* `
* `base64Encode* `
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
* `readDir`
* `readFile`
* `replace`
* `replaceRE`
* `repeat`
* `sha1`
* `sha256`
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

__Usage:__

Import the package:

```
import "github.com/drone/funcmap"
```

Provide the Funcmap when compiling your templates:

```
t, err := template.New("_").Funcs(funcmap.Funcs).Parse(text)
```
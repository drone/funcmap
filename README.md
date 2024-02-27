This package aims to provide an independent implementation of the Hugo template functions.
https://gohugo.io/functions/ ([view source](https://github.com/gohugoio/hugo/tree/7f47b99ea9a7ba7759c4f9424fedd4591e6da497/tpl))

## Usage

Import the package:

```
import "github.com/drone/funcmap"
```

Provide the Funcmap when compiling your templates:

```
t, err := template.New("_").Funcs(funcmap.Funcs).Parse(text)
```

## Functions

### Crypto Functions

* [`md5`](https://gohugo.io/functions/crypto/md5/)
* [`sha1`](https://gohugo.io/functions/crypto/sha1/)
* [`sha256`](https://gohugo.io/functions/crypto/sha256/)

### Encoding Functions

* [`base64Decode`](https://gohugo.io/functions/encoding/base64decode/)
* [`base64Encode`](https://gohugo.io/functions/encoding/base64encode/)
* [`jsonify`](https://gohugo.io/functions/encoding/jsonify/)
* [`jsonEncode`](https://gohugo.io/functions/encoding/jsonify/)
* `jsonDecode`
* `yamlEncode`
* `yamlDecode`

### Date Functions

* [`dateFormat`](https://gohugo.io/functions/time/format/)
* [`now`](https://gohugo.io/functions/time/now/)
* [`time`](https://gohugo.io/functions/time/astime/)

### OS Functions

* [`fileExists`](https://gohugo.io/functions/os/fileexists/)
* [`getenv`](https://gohugo.io/functions/os/getenv/)
* [`readDir`](https://gohugo.io/functions/os/readdir/)
* [`readFile`](https://gohugo.io/functions/os/readfile/)
* [`stat`](https://gohugo.io/functions/os/stat/)

### String Functions

* `append`
* [`chomp`](https://gohugo.io/functions/strings/chomp/)
* [`contains`](https://gohugo.io/functions/strings/contains/)
* [`containsAny`](https://gohugo.io/functions/strings/containsany/)
* [`findRE`](https://gohugo.io/functions/strings/findre/)
* [`firstUpper`](https://gohugo.io/functions/strings/firstupper/)
* [`hasPrefix`](https://gohugo.io/functions/strings/hasprefix/)
* [`hasSuffix`](https://gohugo.io/functions/strings/hassuffix/)
* [`lower`](https://gohugo.io/functions/strings/tolower/)
* `padLeft`
* `padRight`
* `prepend`
* [`repeat`](https://gohugo.io/functions/strings/repeat/)
* [`replace`](https://gohugo.io/functions/strings/replace/)
* [`replaceRE`](https://gohugo.io/functions/strings/replacere/)
* [`slicestr`](https://gohugo.io/functions/strings/slicestring/)
* [`split`](https://gohugo.io/functions/strings/split/)
* `splitn`
* [`title`](https://gohugo.io/functions/strings/title/)
* [`trim`](https://gohugo.io/functions/strings/trim/)
* [`trimLeft`](https://gohugo.io/functions/strings/trimleft/)
* [`trimRight`](https://gohugo.io/functions/strings/trimright/)
* [`trimPrefix`](https://gohugo.io/functions/strings/trimprefix/)
* [`trimSuffix`](https://gohugo.io/functions/strings/trimsuffix/)
* [`upper`](https://gohugo.io/functions/strings/toupper/)

### Transform Functions

* [`htmlEscape`](https://gohugo.io/functions/transform/htmlescape/)
* [`htmlUnescape`](https://gohugo.io/functions/transform/htmlunescape/)

### URL Functions

* [`urlize`](https://gohugo.io/functions/urls/urlize/)

# Background

This is a simple helper function to leverage power of  https://github.com/nicksnyder/go-i18n.

## Usage

Init the locale texts with go:embed :

```go
package main

//go:embed "locale"
var fs embed.FS

func main() {
    i18n.MustInit("locale", &fs)
    // working code
}
```

A typical usage in a web application looks like:

```go
package web

import (
"github.com/wikisio/i18n/i18n"
)

func WebHandler(w http.ResponseWriter, r *http.Request) {
// working code

value := i18n.Get("errMsg001", r.Header.Get("Accept-Language"));

// use the value
}
```

package main

import (
	"embed"
	"fmt"

	"github.com/wikisio/i18n/i18n"
)

//go:embed "locale"
var fs embed.FS

func main() {
	i18n.MustInit("locale", &fs)
	value := i18n.Get("helloId", "zh", "fr,zh-CN;q=0.7,en;q=0.8")
	fmt.Println(value)
}

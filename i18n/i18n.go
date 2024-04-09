package i18n

import (
	"embed"
	"encoding/json"
	"path"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var bundle *i18n.Bundle

// MustInit loads all locale files from the given directory and panics if an error occurs.
func MustInit(localePath string, fs *embed.FS) {

	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	mustLoad(localePath, fs)
}

func mustLoad(localePath string, fs *embed.FS) {
	info, err := fs.ReadDir(localePath)
	if err != nil {
		panic(err)
	}

	for _, f := range info {
		if f.IsDir() {
			mustLoad(path.Join(localePath, f.Name()), fs)
		}

		buf, err := fs.ReadFile(path.Join(localePath, f.Name()))
		if err != nil {
			panic(err)
		}

		bundle.MustParseMessageFileBytes(buf, path.Join(localePath, f.Name()))
	}
}

// Get returns a localized message for the given key and language.
//
// If no language is provided, it will use the default language.
// If the key does not exist in the translation file, it will return the key itself.
func Get(key string, lng ...string) string {
	l := i18n.NewLocalizer(bundle, lng...)
	value, err := l.Localize(&i18n.LocalizeConfig{MessageID: key})
	if err != nil {
		return key
	}

	return value
}

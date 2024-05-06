package config

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed configs/language_extensions.json
var languageExtensionsJSON []byte

type languageExtensions struct {
	Name         string   `json:"name"`
	LanguageType string   `json:"type"`
	Extensions   []string `json:"extensions"`
}

func LanguageExtensionsToMap() (map[string][]string, error) {
	var langsExtensions []languageExtensions
	err := json.Unmarshal(languageExtensionsJSON, &langsExtensions)
	if err != nil {
		return nil, err
	}
	langToExtensions := make(map[string][]string)
	for _, lang := range langsExtensions {
		langToExtensions[strings.ToLower(lang.Name)] = lang.Extensions
	}
	return langToExtensions, nil
}

func FreeLanguageExtensions() {
	languageExtensionsJSON = languageExtensionsJSON[:0]
}

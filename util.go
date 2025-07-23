package flago

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToCamelCase(input string) string {
	words := strings.Split(input, "_")
	caser := cases.Title(language.English)

	for index, word := range words {
		words[index] = caser.String(word)
	}

	return strings.Join(words, "")
}

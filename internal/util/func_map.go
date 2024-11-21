package util

import (
	"net/url"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToCamelCase(s string) string {
	words := strings.Split(s, "_")

	caser := cases.Title(language.Und)
	for i := range words {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = caser.String(words[i])
		}
	}

	return strings.Join(words, "")
}

func ToPascalCase(s string) string {
	words := strings.Split(s, "_")

	caser := cases.Title(language.Und)
	for i := range words {
		words[i] = caser.String(words[i])
	}

	return strings.Join(words, "")
}

func ToLowerCase(s string) string {
	words := strings.Split(s, "_")

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	return strings.Join(words, "")
}

func PathToPascal(s string) string {
	s = strings.Trim(s, "{}") // Remove curly braces if present
	s = strings.ReplaceAll(s, "-", "")
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

func FirstAlphabet(s string) string {
	if len(s) > 0 {
		return string(s[0])
	}
	return s
}

func FirstAlphabetToUpperCase(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

func ExtractPath(s string) string {
	parsedUrl, err := url.Parse(s)
	if err != nil {
		return ""
	}

	return parsedUrl.Path
}

func JoinStrings(sep string, items []string) string {
	return strings.Join(items, sep)
}

func CreateFuncMap() template.FuncMap {
	return template.FuncMap{
		"ToCamelCase":              ToCamelCase,
		"ToPascalCase":             ToPascalCase,
		"ToLowerCase":              ToLowerCase,
		"PathToPascal":             PathToPascal,
		"FirstAlphabet":            FirstAlphabet,
		"FirstAlphabetToUpperCase": FirstAlphabetToUpperCase,
		"JoinStrings":              JoinStrings,
		"ExtractPath":              ExtractPath,
	}
}

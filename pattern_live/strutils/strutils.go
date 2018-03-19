package strutils

import (
	"strings"
	"unicode"
)

func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

func ToFirstUpper(s string) string {
	if len(s) < 1 {
		return s
	}
	// Trim the string
	t := strings.Trim(s, " ")
	// Convert all letters to lower case
	t = strings.ToLower(t)
	res := []rune(t)
	// Convert first letter to upper case
	res[0] = unicode.ToUpper(res[0])
	return string(res)
}


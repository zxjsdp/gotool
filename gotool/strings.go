package gotool

import (
	"strings"
	"unicode"
)

// RemoveChar removes rune from string.
func RemoveRune(source string, runeToBeRemoved rune) string {
	return strings.Map(func(r rune) rune {
		if r != runeToBeRemoved {
			return r
		}
		return -1
	}, source)
}

func RemoveBlankRunes(source string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsSpace(r) {
			return r
		}
		return -1
	}, source)
}

// ReplaceRunes replaces blank chars in string with new rune.
func ReplaceBlankRunes(source string, newRune rune) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return newRune
		}
		return r
	}, source)
}

// ReplaceRunes replaces blank chars in string with new token.
func ReplaceRunes(source string, oldRune, newRune rune) string {
	return strings.Map(func(r rune) rune {
		if r == oldRune {
			return newRune
		}
		return r
	}, source)
}

// CheckSubStringExistence checks whether substring exists in string.
func CheckSubStringExistence(str, subString string) bool {
	index := strings.Index(str, subString)
	if index >= 0 {
		return true
	}
	return false
}

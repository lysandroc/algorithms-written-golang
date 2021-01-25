// Package isogram contains a function to validate if a word is an isogram
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(word string) bool {
	isogramMap := make(map[rune]int)

	for _, char := range strings.ToUpper(word) {
		isogramMap[char]++
	}

	for char, count := range isogramMap {
		if count > 1 && unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

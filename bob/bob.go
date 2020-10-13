package bob

import (
	"strings"
	"unicode"
)

const blankSpace string = "\n\r \t"

//Hey returns a limited Bob's answer
func Hey(remark string) string {
	trimWithoutWhitespace := strings.Trim(remark, blankSpace)
	isRemarkEmpty := trimWithoutWhitespace == ""
	isAQuestion := strings.HasSuffix(trimWithoutWhitespace, "?")

	isAllUpperCase := isWordUpperCase(remark)

	if isAQuestion && isAllUpperCase {
		return "Calm down, I know what I'm doing!"
	} else if isAllUpperCase {
		return "Whoa, chill out!"
	} else if isAQuestion {
		return "Sure."
	} else if isRemarkEmpty {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func isWordUpperCase(word string) bool {
	countLetters := 0
	countLettersIsUpperCase := 0

	for _, letter := range word {
		isLetter := unicode.IsLetter(letter)
		isUpper := unicode.IsUpper(letter)
		if isLetter {
			countLetters++
			if isUpper {
				countLettersIsUpperCase++
			}
		}
	}

	if countLetters == 0 {
		return false
	}
	return countLetters == countLettersIsUpperCase
}

package scrabble

import "strings"

//Given the same structure
var valueByLetters = map[int][]string{
	1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  {"D", "G"},
	3:  {"B", "C", "M", "P"},
	4:  {"F", "H", "V", "W", "Y"},
	5:  {"K"},
	8:  {"J", "X"},
	10: {"Q", "Z"},
}

// transformDataFormat receives the original format structure then it returns the letter and point pairs.
func transformDataFormat(valueByLetters map[int][]string) map[string]int {
	valueByLetter := make(map[string]int)
	for point, letters := range valueByLetters {
		for _, letter := range letters {
			valueByLetter[letter] = point
		}
	}
	return valueByLetter
}

// Score calculates the Scrabble value for the word received
func Score(word string) (score int) {
	for _, r := range word {
		for letter, points := range transformDataFormat(valueByLetters) {
			if strings.ToUpper(string(r)) == strings.ToUpper(string(letter)) {
				score += points
			}
		}
	}
	return
}

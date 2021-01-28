package luhn

import (
	"strconv"
	"strings"
)

// treatCardNumber returns the check digit and the rest of card number reversed
func treatCardNumber(cardNumber string) (string, string) {
	reverseCardNumber := ""
	for _, v := range cardNumber {
		reverseCardNumber = string(v) + reverseCardNumber
	}

	return reverseCardNumber[:1], reverseCardNumber[1:]
}

// Valid verify wether card number is valid or not
func Valid(cardNumber string) bool {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	cardNumber = strings.TrimSpace(cardNumber)

	if len(cardNumber) <= 1 {
		return false
	}

	checkDigitChar, digits := treatCardNumber(cardNumber)

	checkDigit, err := strconv.Atoi(string(checkDigitChar))
	if err != nil {
		return false
	}

	sum := 0
	for i, digit := range digits {
		digit, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}

		if i%2 == 0 {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return (sum+checkDigit)%10 == 0
}

package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

// splitCardNumber returns the check digit and the rest of card number reversed
func treatCardNumber(cardNumber string) (checkDigit int, reverseCardNumber string, err error) {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	cardNumber = strings.TrimSpace(cardNumber)

	if len(cardNumber) <= 1 {
		return 0, "", fmt.Errorf("the card number length is invalid!")
	}

	for _, v := range cardNumber {
		reverseCardNumber = string(v) + reverseCardNumber
	}

	digitCheck, err := strconv.Atoi(string(reverseCardNumber[:1]))

	return digitCheck, reverseCardNumber[1:], err
}

// Valid verify wether card number is valid or not
func Valid(cardNumber string) bool {
	var sum int

	checkDigit, digits, err := treatCardNumber(cardNumber)
	if err != nil {
		return false
	}

	for i, digit := range digits {
		digit, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}

		if isEven := i%2 == 0; isEven {
			if digit = digit * 2; digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return (sum+checkDigit)%10 == 0
}

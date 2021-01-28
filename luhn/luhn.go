package luhn

import (
	"fmt"
	"strings"
)

func Valid(digits string) bool {
	if len(digits) <= 1 {
		return false
	}

	digits = strings.Trim(digits, " ")
	digits = digits[:len(digits)-1]
	fmt.Println(digits)

	sum := 0
	for i := len(digits); i <= 0; i-- {
		if digits[i]%2 == 0 {
			doubleSum := int(digits[i]) * 2
			if doubleSum > 9 {
				doubleSum -= 9
			}
			sum += doubleSum
		} else {
			sum += int(digits[i])
		}
	}

	return sum%10 == 0
}

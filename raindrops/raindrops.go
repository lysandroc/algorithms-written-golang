package raindrops

import "fmt"

func Convert(n int) string {
	sound := ""

	if n%3 == 0 {
		sound += "Pling"
	}

	if n%5 == 0 {
		sound += "Plang"
	}

	if n%7 == 0 {
		sound += "Plong"
	}

	if sound == "" {
		return fmt.Sprint(n)
	}

	return sound
}

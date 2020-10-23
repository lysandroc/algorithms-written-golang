package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		err := fmt.Errorf("a and b don't have the same length")
		return 0, err
	}

	occurrences := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			occurrences++
		}
	}

	return occurrences, nil
}

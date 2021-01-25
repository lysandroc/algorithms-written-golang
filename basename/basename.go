package basename

import (
	"strings"
)

func BasenameV1(s string) string {
	// Ignore the last "/" and all characters before it
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	//Prevents all before the last "."
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func Basename(s string) string {
	slash := strings.LastIndex(s, "/") // returns -1 if it doesn't exist
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//twofer package contains a method that returns a 2-fer string
package twofer

import "fmt"

//ShareWith receives a name and returns a 2-fer message
func ShareWith(name string) string {
	personName := name
	if name == "" {
		personName = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", personName)
}

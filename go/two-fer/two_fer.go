// twofer package helps to share with people
package twofer

import "fmt"

// ShareWith returns share phrase including passed name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}

package cli

import (
	"fmt"
)

// Author one who writes a book
type Author struct {
	Alias string
	firstName string
	lastName string
	age int
}

// FullName returns the last and first name
func (a *Author) FullName() string {
	return fmt.Sprintf("%s %s", a.lastName, a.firstName)
}

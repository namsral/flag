package flag

import (
	"flag"
	"fmt"
	"strings"
)

// Checker is an interface for type which has a Check function
type Checker interface {
	Check() error
}

// Validate runs flag checkers
func Validate(positionalArgs string, checkers ...Checker) error {
	expected := strings.Fields(positionalArgs)
	if len(expected) != flag.NArg() {
		return fmt.Errorf("missing required positional arguments: %s. Number of args Provided: %d, expected: %d",
			expected, flag.NArg(), len(expected))
	}
	for _, c := range checkers {
		if err := c.Check(); err != nil {
			return err
		}
	}
	return nil
}

// CheckMany is a helper function to check many flag components
func CheckMany(checkers ...Checker) error {
	for _, c := range checkers {
		if err := c.Check(); err != nil {
			return err
		}
	}
	return nil
}

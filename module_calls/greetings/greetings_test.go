package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, and check
// for valid return value
func TestHelloName(t *testing.T) {
	name := "GladOS"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("GladOS" = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hell with empty string,
// expects an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(
			`Hello("") = %q, %v, want "", error`,
			msg,
			err,
		)
	}
}

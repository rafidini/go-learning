package api

import (
    "testing"
)


func TestSuccessJohnDoe(t *testing.T) {
    var actual   string = JohnDoe()
	var expected string = "John Doe"

    if actual != expected {
        t.Fatalf(`JohnDoe() = %q, want match for %#q`, actual, expected)
    }
}

func TestFailJohnDoe(t *testing.T) {
    var actual   string = JohnDoe()
	var expected string = "Mamo"

    if actual == expected {
        t.Fatalf(`JohnDoe() = %q, want match for %#q`, actual, expected)
    }
}

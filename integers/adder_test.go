package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// This code will be documented with godoc and will be shown as an example.
// Don't omit the Output line, it allows the parser to show the output in another section.
// run godoc -http :8080 and visit the port to the docs.
func ExampleAdd() {
	sum := Add(1, 4)
	fmt.Println(sum)
	// Output: 5
}

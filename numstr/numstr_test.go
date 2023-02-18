package numstr_test

import (
	"fmt"

	"github.com/pschou/go-numstr"
)

func ExampleLessThanFold() {
	a, b := "abc123", "abc12"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	a, b = "abc100", "abc121"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	a, b = "abc10a", "abc1d"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	// Output:
	// abc123 < abc12  false
	// abc100 < abc121  true
	// abc10a < abc1d  false
}

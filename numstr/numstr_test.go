package numstr_test

import (
	"fmt"

	"github.com/pschou/go-sorting/numstr"
)

func ExampleLessThanFold() {
	a, b := "abc123", "abc12"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	a, b = "abc100", "abc121"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	a, b = "abc10a", "abc1d"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	a, b = "abc.txt", "abc1.txt"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	a, b = "abc1.txt", "abc1a.txt"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	a, b = "a01", "a0a"
	fmt.Printf("%s < %s  %v\n", a, b, numstr.LessThanFold(a, b))
	// Output:
	// abc123 < abc12  false
	// abc100 < abc121  true
	// abc10a < abc1d  false
	// abc.txt < abc1.txt  true
	// abc1.txt < abc1a.txt  true
}

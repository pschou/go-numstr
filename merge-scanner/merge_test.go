package merge

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func ExampleMerge() {
	A := `aa
ba
ca`
	B := `ab
ac
af
bas
bf`
	C := `ae
ag
baaad
bg`
	Ar := bufio.NewScanner(strings.NewReader(A))
	Br := bufio.NewScanner(strings.NewReader(B))
	Cr := bufio.NewScanner(strings.NewReader(C))
	Mr := NewWithCompare(func(a, b []byte) int {
		return bytes.Compare(a[:2], b[:2])
	}, Ar, Br, Cr)
	for Mr.Scan() {
		fmt.Println(Mr.Text())
	}
	// Output:
	// aa
	// ab
	// ac
	// ae
	// af
	// ag
	// ba
	// bf
	// bg
	// ca
}

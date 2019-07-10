package main

import (
	"fmt"
)

func main() {
	n := 1001 * 1011
	s := fmt.Sprintf("%d", n)
	if s[0] == s[len(s)-1] {
		fmt.Printf("%q == %q", s[0], s[len(s)-1])
		fmt.Printf("\nthe number is even-ended  %q", s)
	}

}

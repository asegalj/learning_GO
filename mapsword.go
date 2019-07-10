package main

import (
	"fmt"
	"strings"
)

func main() {
	wordMapper := map[string]int{}
	text := `Needles and pins 
			 Needles and Pins
			 sew me Sail
			 tO catCH me the wind`
	b := strings.ToLower(text)
	a := strings.Fields(b)
	fmt.Println(a)
	for _, value := range a {
		// fmt.Println(wordMapper[value]=wordMapper[value]+1)
		wordMapper[value] = wordMapper[value] + 1
	}
	fmt.Println(wordMapper)
}

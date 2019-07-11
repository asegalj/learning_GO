package main

import (
	"fmt"
	"strings"
)

//My solution
// func main() {
// 	wordMapper := map[string]int{}
// 	text := `Needles and pins
// 			 Needles and Pins
// 			 sew me Sail
// 			 tO catCH me the wind`
// 	b := strings.ToLower(text)
// 	a := strings.Fields(b)
// 	fmt.Println(a)
// 	for _, value := range a {
// 		// fmt.Println(wordMapper[value]=wordMapper[value]+1)
// 		wordMapper[value]++
// 	}
// 	fmt.Println(wordMapper)
// }

func main() {

	text := `Needles and pins 
			 Needles and Pins
			 sew me Sail
			 tO catCH me the wind`
	words := strings.Fields(text)
	wordMapper := map[string]int{}
	for _, word := range words {
		// fmt.Println(wordMapper[value]=wordMapper[value]+1)
		wordMapper[strings.ToLower(word)]++
	}
	fmt.Println(wordMapper)
}

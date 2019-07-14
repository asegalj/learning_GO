package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d \n", div, mod)

	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values)

	val1 := 10
	double(val1)
	fmt.Println(val1)
	val2 := 10
	doublePtr(&val2)
	fmt.Println(val2)
}

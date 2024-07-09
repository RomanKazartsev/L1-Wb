package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	// Поменять местами значения a и b
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a = %d, b = %d\n", a, b) // Выведет: a = 10, b = 5
}

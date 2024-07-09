package main

import (
	"fmt"
)

func main() {
	var i interface{} = 44
	switch v := i.(type) {
	case int:
		fmt.Printf("Переменная имеет тип int: %d\n", v)
	case string:
		fmt.Printf("Переменная имеет тип string: %q\n", v)
	case bool:
		fmt.Printf("Переменная имеет тип bool: %t\n", v)
	case chan int:
		fmt.Printf("Переменная имеет тип chan int: %v\n", v)
	default:
		fmt.Printf("Неизвестный тип: %T\n", v)
	}
}

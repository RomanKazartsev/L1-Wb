package main

import (
	"fmt"
)

// ReverseString переворачивает строку, поддерживая Unicode символы
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	inputString := "главрыба"
	reversedString := ReverseString(inputString)
	fmt.Printf("Оригинальная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}


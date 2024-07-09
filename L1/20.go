package main

import (
	"fmt"
	"strings"
)

// ReverseWords переворачивает слова в строке
func ReverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	inputString := "snow dog sun"
	reversedWords := ReverseWords(inputString)
	fmt.Printf("Оригинальная строка: %s\n", inputString)
	fmt.Printf("Строка с перевернутыми словами: %s\n", reversedWords)
}

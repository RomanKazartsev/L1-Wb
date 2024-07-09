package main

import (
	"fmt"
	"strings"
)

// IsUnique проверяет, все ли символы в строке уникальны
func IsUnique(str string) bool {
	charMap := make(map[rune]bool)
	for _, char := range strings.ToLower(str) {
		if _, exists := charMap[char]; exists {
			return false // Найден повторяющийся символ
		}
		charMap[char] = true
	}
	return true // Все символы уникальны
}

func main() {
	testStrings := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, str := range testStrings {
		fmt.Printf("Строка \"%s\" содержит только уникальные символы: %t\n", str, IsUnique(str))
	}
}

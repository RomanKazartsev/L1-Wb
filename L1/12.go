package main

import (
	"fmt"
)

// Функция для создания множества из среза строк
func createSet(strings []string) []string {
	setMap := make(map[string]bool)
	uniqueSet := []string{}

	// Добавление элементов в карту для обеспечения уникальности
	for _, value := range strings {
		if _, exists := setMap[value]; !exists {
			setMap[value] = true
			uniqueSet = append(uniqueSet, value)
		}
	}

	return uniqueSet
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	// Вызов функции и вывод результата
	fmt.Println(createSet(strings)) // Выведет: [cat dog tree]
}

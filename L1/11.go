package main

import (
	"fmt"
)

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 []int) []int {
	set1Map := make(map[int]bool)
	intersectionSet := []int{}

	// Заполнение карты элементами первого множества
	for _, value := range set1 {
		set1Map[value] = true
	}

	// Поиск общих элементов во втором множестве
	for _, value := range set2 {
		if _, found := set1Map[value]; found {
			intersectionSet = append(intersectionSet, value)
		}
	}

	return intersectionSet
}

func main() {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{4, 5, 6, 7, 8}

	// Вызов функции и вывод результата
	fmt.Println(intersection(list1, list2)) // Выведет: [4 5]
}

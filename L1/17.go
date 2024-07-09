package main

import (
	"fmt"
	"sort"
)

func binarySearch(slice []int, target int) int {
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= target
	})
	if index < len(slice) && slice[index] == target {
		return index // Элемент найден
	}
	return -1 // Элемент не найден
}

func main() {
	slice := []int{1, 2, 4, 5, 6, 8, 9}
	target := 5
	result := binarySearch(slice, target)
	if result != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, result)
	} else {
		fmt.Println("Элемент не найден.")
	}
}


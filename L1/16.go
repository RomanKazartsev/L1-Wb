package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Исходный массив:", arr)
	sort.Ints(arr)
	fmt.Println("Отсортированный массив:", arr)
}

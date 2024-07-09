package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	// Функция для определения ключа группы
	getGroupKey := func(temp float64) int {
		return int(math.Floor(temp/10)) * 10
	}

	// Группировка температур
	for _, temp := range temperatures {
		key := getGroupKey(temp)
		groups[key] = append(groups[key], temp)
	}

	// Вывод результатов
	for key, temps := range groups {
		fmt.Printf("%d:{%v}\n", key, temps)
	}
}

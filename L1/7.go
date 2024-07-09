package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	// Инициализация sync.Map
	var concurrentMap sync.Map

	// Функция для записи данных в map
	writeData := func(key, value int) {
		defer wg.Done()
		concurrentMap.Store(key, value)
		fmt.Printf("Записано: ключ %v, значение %v\n", key, value)
	}

	// Запускаем несколько горутин для записи данных в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeData(i, i*i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим данные из map
	concurrentMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Ключ: %v, Значение: %v\n", key, value)
		return true // Продолжаем итерацию по map
	})
}

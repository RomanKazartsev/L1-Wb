


package main

import (
	"fmt"
	"sync"
)

// Counter - структура для счетчика
type Counter struct {
	mu    sync.Mutex // Мьютекс для синхронизации
	value int        // Значение счетчика
}

// Increment увеличивает значение счетчика на 1
func (c *Counter) Increment() {
	c.mu.Lock()         // Блокируем доступ к value
	c.value++           // Инкрементируем счетчик
	c.mu.Unlock()       // Разблокируем доступ
}

// Value возвращает текущее значение счетчика
func (c *Counter) Value() int {
	c.mu.Lock()         // Блокируем доступ к value
	defer c.mu.Unlock() // Отложенная разблокировка доступа
	return c.value      // Возвращаем значение счетчика
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Количество горутин для запуска
	numGoroutines := 1000

	// Добавляем количество ожидаемых горутин в WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done() // Уменьшаем счетчик WaitGroup по завершении горутины
			counter.Increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Итоговое значение счетчика: %d\n", counter.Value())
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10} // Исходный массив чисел
	squares := make(chan int)        // Канал для сбора квадратов чисел
	var wg sync.WaitGroup            // WaitGroup для ожидания завершения горутин
	sum := 0                         // Переменная для суммы квадратов

	// Функция для вычисления квадрата числа и отправки результата в канал
	calcSquare := func(n int) {
		defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения горутины
		squares <- n * n
	}

	// Создаем горутину для каждого числа в массиве
	for _, number := range numbers {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go calcSquare(number)
	}

	// Запускаем горутину для сбора результатов из канала
	go func() {
		for square := range squares {
			sum += square
		}
	}()

	// Ожидаем завершения всех горутин
	wg.Wait()
	close(squares) // Закрываем канал после завершения всех горутин

	// Выводим сумму квадратов
	fmt.Printf("Сумма квадратов чисел: %d\n", sum)
}

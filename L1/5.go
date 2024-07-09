package main

import (
	"fmt"
	"time"
)

func main() {
	// Задаем количество секунд для таймера
	var N int
	fmt.Print("Введите количество секунд N: ")
	fmt.Scan(&N)

	// Создаем канал для передачи целых чисел
	dataChannel := make(chan int)

	// Запускаем горутину для отправки значений в канал
	go func() {
		for i := 0; ; i++ {
			dataChannel <- i
			time.Sleep(1 * time.Second) // Имитация работы
		}
	}()

	// Создаем таймер для завершения программы
	timer := time.NewTimer(time.Duration(N) * time.Second)

	// Читаем значения из канала и обрабатываем таймер
	for {
		select {
		case value := <-dataChannel:
			fmt.Println("Получено значение:", value)
		case <-timer.C:
			fmt.Println("Время вышло, программа завершается...")
			close(dataChannel)
			return
		}
	}
}

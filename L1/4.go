package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// Проверяем, передано ли количество воркеров как аргумент командной строки
	if len(os.Args) < 2 {
		fmt.Println("Необходимо указать количество воркеров")
		return
	}

	// Конвертируем аргумент командной строки в число
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Ошибка при конвертации количества воркеров:", err)
		return
	}

	// Создаем канал для передачи данных
	dataChannel := make(chan int)
	// Канал для обработки сигналов
	signals := make(chan os.Signal, 1)
	// Регистрируем сигналы
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Функция воркера
	worker := func(id int, data <-chan int, done <-chan struct{}) {
		for {
			select {
			case d, ok := <-data:
				if !ok {
					fmt.Printf("Воркер %d завершает работу\n", id)
					return
				}
				fmt.Printf("Воркер %d получил данные: %d\n", id, d)
			case <-done:
				fmt.Printf("Воркер %d получил сигнал на завершение\n", id)
				return
			}
		}
	}

	// Канал для оповещения воркеров о завершении работы
	done := make(chan struct{})

	// Запускаем N воркеров
	for i := 0; i < n; i++ {
		go worker(i, dataChannel, done)
	}

	// Генерируем данные и записываем их в канал
	go func() {
		for {
			select {
			case <-signals:
				return
			default:
				dataChannel <- rand.Intn(100)
				time.Sleep(time.Second)
			}
		}
	}()

	// Ожидаем сигнала
	<-signals
	close(done)     // Оповещаем воркеров о завершении работы
	close(dataChannel) // Закрываем канал с данными
	fmt.Println("Программа завершает работу")
}

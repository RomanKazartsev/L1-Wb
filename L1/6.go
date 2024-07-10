package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Использование каналов для сигнализации
func workerWithChannel(stopChan <-chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Горутина остановлена через канал")
			return
		default:
			fmt.Println("Горутина работает (канал)")
			time.Sleep(1 * time.Second)
		}
	}
}

// Использование контекста с отменой
func workerWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановлена через контекст:", ctx.Err())
			return
		default:
			fmt.Println("Горутина работает (контекст)")
			time.Sleep(1 * time.Second)
		}
	}
}

// Использование WaitGroup
func workerWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина %d начала работу (WaitGroup)\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Горутина %d завершила работу (WaitGroup)\n", id)
}

func main() {
	// Каналы
	stopChan := make(chan struct{})
	go workerWithChannel(stopChan)

	// Контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	go workerWithContext(ctx)

	// WaitGroup
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerWithWaitGroup(i, &wg)
	}

	// Остановка горутины через каналы после 3 секунд
	time.Sleep(3 * time.Second)
	close(stopChan)

	// Остановка горутины через контекст после 4 секунд
	time.Sleep(1 * time.Second)
	cancel()

	// Ожидание завершения горутин через WaitGroup
	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}

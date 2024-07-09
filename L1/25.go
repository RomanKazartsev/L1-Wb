package main

import (
	"fmt"
	"time"
)

// Sleep функция, которая приостанавливает выполнение на заданное количество секунд
func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало сна...")
	Sleep(2) // Спим 2 секунды
	fmt.Println("Конец сна.")
}

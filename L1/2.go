package main

import (
    "fmt"
    "sync"
)

func main() {
    // Массив чисел для расчета квадратов
    numbers := []int{2, 4, 6, 8, 10}
    // WaitGroup для ожидания завершения всех горутин
    var wg sync.WaitGroup

    // Функция для расчета квадрата числа и вывода результата
    square := func(n int) {
        defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения горутины
        fmt.Printf("Квадрат числа %d равен %d\n", n, n*n)
    }

    // Создаем горутину для каждого числа в массиве
    for _, number := range numbers {
        wg.Add(1) // Увеличиваем счетчик WaitGroup
        go square(number)
    }

    // Ожидаем завершения всех горутин
    wg.Wait()
}
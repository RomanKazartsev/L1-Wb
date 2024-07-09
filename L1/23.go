package main

import "fmt"

func main() {
	// Задаем слайс
	slice := []int{1, 2, 3, 4, 5}
	// Указываем индекс элемента для удаления
	i := 2 // индекс элемента, который нужно удалить

	// Удаление i-го элемента
	slice = append(slice[:i], slice[i+1:]...)

	fmt.Println(slice)
}

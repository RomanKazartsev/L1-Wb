package main

import (
	"fmt"
)

func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

func clearBit(n int64, pos uint) int64 {
	mask := ^(1 << pos)
	n &= int64(mask)
	return n
}

func main() {
	var num int64 = 0 // Исходное число
	var i uint = 5    // Позиция бита, который нужно изменить

	// Установка i-го бита в 1
	num = setBit(num, i)
	fmt.Printf("Число с установленным битом: %b\n", num)

	// Установка i-го бита в 0
	num = clearBit(num, i)
	fmt.Printf("Число с очищенным битом: %b\n", num)
}
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация переменных a и b значениями больше 2^20
	a := new(big.Int).SetInt64(1 << 21) // a = 2^21
	b := new(big.Int).SetInt64(1 << 22) // b = 2^22

	// Умножение
	mul := new(big.Int).Mul(a, b)
	fmt.Println("Умножение:", mul)

	// Деление
	div := new(big.Int).Quo(b, a)
	fmt.Println("Деление:", div)

	// Сложение
	add := new(big.Int).Add(a, b)
	fmt.Println("Сложение:", add)

	// Вычитание
	sub := new(big.Int).Sub(a, b)
	fmt.Println("Вычитание:", sub)
}


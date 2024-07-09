package main

import (
	"fmt"
	"math"
)

// Point - структура для представления точки в 2D пространстве
type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Distance - метод для вычисления расстояния между двумя точками
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	// Создаем две точки с помощью конструктора
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	// Вычисляем и выводим расстояние между точками
	fmt.Printf("Расстояние между точками: %.2f\n", p1.Distance(p2))
}

package main

import "fmt"

// Human - родительская структура с набором полей и методов
type Human struct {
    Name string
    Age  int
}

// Speak - метод структуры Human
func (h *Human) Speak() {
    fmt.Println("Меня зовут", h.Name)
}

// Action - дочерняя структура, встраивающая Human
type Action struct {
    Human // Встраивание структуры Human
    ActionDescription string
}

func main() {
    // Создание экземпляра структуры Action
    action := Action{
        Human: Human{
            Name: "Алексей",
            Age:  30,
        },
        ActionDescription: "бегает по утрам",
    }

    // Вызов метода Speak от встроенной структуры Human
    action.Speak()
    fmt.Println("и", action.ActionDescription)
}

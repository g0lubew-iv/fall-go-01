// Package personaldata содержит структуры и методы для работы с личными данными пользователя
package personaldata

import "fmt"

// Personal хранит базовую информацию о пользователе: имя, вес и рост
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит личные данные пользователя в стандартный поток вывода
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}

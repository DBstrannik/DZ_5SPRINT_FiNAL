package personaldata

import "fmt"

// Personal описывает персональные данные пользователя.
type Personal struct {
	Name   string  // имя пользователя.
	Weight float64 // вес пользователя
	Height float64 // рост пользователя

}

// Print выводит персональьные данные пользователя на экран
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.1f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}

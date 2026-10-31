package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Println("Имя:", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}

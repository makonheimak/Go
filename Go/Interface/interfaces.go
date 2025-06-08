package main

import "fmt"

func main() {
	move := []Movable{
		Car{Brand: "Bugatti_Leturbini", Speed: "1450"},
		Person{Name: "Oleg", Age: "35", Speed: "5"},
	}

	for _, m := range move {
		printMovement(m)
	}

}

type Car struct {
	Brand string
	Speed string
}

func (c Car) Move() string {
	return "Машина едет со скоростью " + c.Speed + " км/ч"
}

type Movable interface {
	Move() string
}

type Person struct {
	Name  string
	Age   string
	Speed string
}

func (p Person) Move() string {
	return "Человек идёт со скоростью " + p.Speed + " км/ч"
}

func printMovement(m Movable) {
	fmt.Println(m.Move())
}

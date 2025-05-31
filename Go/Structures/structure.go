package main

import "fmt"

func main() {
	myCar := Car{
		Brand: "Bugatti_Leturbini",
		Year:  2025,
	}
	fmt.Print(myCar)

	updateCarYear(myCar)
}

type Car struct {
	Brand string
	Year  int
}

func updateCarYear(car Car) {
	car.Year++
	fmt.Println("\nВнутри функции:", car)
}

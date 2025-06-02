package main

import "fmt"

func main() {
	myCar := &Car{Brand: "Bugatti_Leturbini", Year: 2025}
	fmt.Print("До:", myCar.Year)

	updateCarYearPointer(myCar)
	fmt.Println("\nПосле:", myCar.Year)
}

type Car struct {
	Brand string
	Year  int
}

func updateCarYearPointer(car *Car) {
	car.Year++
}

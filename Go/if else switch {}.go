package main

import "fmt"

func main() {
	isSunny := false
	isWeekend := true

	if isSunny && isWeekend == true {
		fmt.Println("Идеальный день для прогулки!")
	} else if isSunny && !isWeekend {
		fmt.Println("Погода хорошая, но нужно поработать :(")
	} else if !isSunny && isWeekend {
		fmt.Println("Можно остаться дома и отдохнуть")
	} else if !isSunny && !isWeekend {
		fmt.Println("Рабочий день с плохой погодой")
	}
}

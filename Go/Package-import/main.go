package main

import (
	"fmt"
	"packageimport/converter"
	mp1 "packageimport/mathpack"
)

func main() {
	c := 38.4
	f := converter.CelsiusToFahrenheit(c)
	fmt.Println("Фаренгейтов:", f)

	km := 12.4
	miles := converter.KilometersToMiles(km)
	fmt.Println("Миль:", miles)

	base := 10
	exponent := 4
	result1 := mp1.Pow(base, exponent)
	fmt.Println("Число в степени: ", result1)

	n := 11
	result2 := mp1.IsPrime(n)
	if result2 {
		fmt.Println("Число ", n, " простое")
	} else {
		fmt.Println("Число ", n, " не простое")
	}

}

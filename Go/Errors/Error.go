package main

import (
	"errors"
	"fmt"
)

var ErrUnderAge = errors.New("возраст меньше нуля")

func main() {

	result, err1 := divide(12, 12)
	if err1 != nil {
		fmt.Println("Ошибка:1")
		return
	}

	fmt.Println("Результат:", result)

	age := 0
	err2 := validateAge(age)
	if errors.Is(err2, ErrUnderAge) {
		fmt.Println("Ошибка:2")
		return
	}
	fmt.Println("Возраст корректен", age)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Ошибка: Деление на 0!")
	}
	return a / b, nil
}

func validateAge(age int) error {
	if age <= 0 {
		return ErrUnderAge
	}
	return nil
}

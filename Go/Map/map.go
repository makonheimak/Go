package main

import "fmt"

func main() {

	menu := map[string]float64{
		"RAF Coffi": 352.50,
		"Сookies":   129.99,
	}

	menu["Цезарь с креветками"] = 634.67

	if _, ok := menu["Салат"]; !ok {
		menu["Салат"] = 12.99
		fmt.Println("Добавлен Салат в меню")
	} else {
		fmt.Println("Салат уже есть в меню!")
	}

	delete(menu, "Сookies")

	for i, v := range menu {
		fmt.Println(i, v)
	}

}

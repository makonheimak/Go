package main

import "fmt"

func main() {
	Laptops := []Laptop{
		{Brand: "Lenivo Legion 5 Pro ", Price: "1499 $ \n"},
		{Brand: "Lenivo Legion 5 ", Price: "1099 $\n"},
		{Brand: "ASUS Tuf Gaming ", Price: "1340 $\n"},
	}
	fmt.Println("Исходный список ноутбуков:\n", Laptops)

	newlaptop := Laptop{
		Brand: "ASUS Gambling Laptop",
		Price: "1999 $",
	}

	Laptops = append(Laptops, newlaptop)
	fmt.Println("Новый ноутбук:\n", Laptops)

	Laptops[1] = Laptop{Brand: "Lenivo Legion 5", Price: "10999 $\n"}
	fmt.Println("Новая цена:\n", Laptops)

	visov(Laptops)
}

type Laptop struct {
	Brand string
	Price string
}

func visov(data []Laptop) {
	for index, laptop := range data {
		fmt.Println("Вывод for:", index, laptop)
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("номер: ", i)
		time.Sleep(500 * time.Millisecond) // 0.5 сек
		if i == 6 {
			fmt.Println("stop")
			break
		}
	}
}

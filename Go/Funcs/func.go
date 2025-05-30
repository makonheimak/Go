/*package main

import "fmt"

func input1(input int) int {
	return input * 2
}
func main() {
	input := 5
	fmt.Println(input1(input))
}
*/

package main

import "fmt"

func input(a, b int) (int, int) {
	input1 := a + b
	input2 := a * b
	return input1, input2
}
func main() {
	a := 6
	b := 10
	input1, input2 := input(a, b)

	fmt.Println("Сумма:", input1)
	fmt.Println("Произведение:", input2)
}

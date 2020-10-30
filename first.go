package main

import "fmt"

func main() {

	a := 1
	b := 2
	res := calc(a, b)
	fmt.Print(res)

}

func calc(num1, num2 int) int {
	var total = 0
	for i := 1; i <= 5; i++ {
		total = num1 + num2

	}
	return total
}

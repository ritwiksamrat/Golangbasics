package main

import "fmt"

func main() {

	ans := add(5, 10)
	fmt.Println(ans)

}

func add(x int, y int) int {
	defer fmt.Println("Two numbers will be add")
	fmt.Println("This will be excute in second")
	return x + y
}

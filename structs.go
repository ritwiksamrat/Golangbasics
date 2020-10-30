package main

import "fmt"

type Rit struct {
	x int64
	y int64
}

func main() {

	var a = Rit{5, 7} 
	var b = Rit{-2, -2}
	mul(a.y)
	fmt.Println(res)

}

func mul(z int) int {

	multi := z * 2
	return multi
}

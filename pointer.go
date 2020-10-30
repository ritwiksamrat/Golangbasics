package main

import "fmt"

func main() {

	change := "hello"
	fmt.Println(change)
	tochange(&change)
	fmt.Println(change)

}

func tochange(str *string) {

	*str = "Its changed"

}

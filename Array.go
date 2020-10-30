package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Enter the size of the array")

	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	var arr [input]int
	for i := 0; i < n; i++ {

		fmt.Scan(&arr[i])
	}

	fmt.Println(arr)

}

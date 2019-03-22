package main

import (
	"fmt"
)

func changeV1(i int) {
	i = 10
}

func changeV2(i *int) {
	*i = 10
}
func main() {
	num := 0

	// var num1 *int = &num

	// changeV1(num)
	changeV2(&num)
	fmt.Println(num)
}

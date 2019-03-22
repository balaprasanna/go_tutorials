package main

import (
	"fmt"
)

func perform(op string, nums ...int) (total int) {
	if op == "*" {
		total = 1
	}
	for _, v := range nums {
		switch op {
		case "+":
			total += v
		case "*":
			total *= v
		default:
			total += v
		}
	}
	return
}

func main() {

	numbers := []int{10, 20, 30}
	fmt.Println("SUM = ", perform("+", numbers...))
	fmt.Println("Prod = ", perform("*", 10, 20, 30, 40))

}

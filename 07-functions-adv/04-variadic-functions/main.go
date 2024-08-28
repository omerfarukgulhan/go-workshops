package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sumOfNumbers1 := sum(10, 15, 40, -5)
	sumOfNumbers2 := sum(numbers...)

	fmt.Println(sumOfNumbers1)
	fmt.Println(sumOfNumbers2)
}

func sum(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

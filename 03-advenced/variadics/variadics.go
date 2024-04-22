package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{11, 22, 33, 44}

	fmt.Println(sum(a...))
	fmt.Println(sum(b...))
	fmt.Println(sum(1, 2, 3, 4, 5))
}

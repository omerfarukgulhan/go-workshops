package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add("a", "b"))
	fmt.Println(add(1.42, 2.512))
	fmt.Println(add(12, 5))
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

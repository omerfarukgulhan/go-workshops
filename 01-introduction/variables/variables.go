package main

import "fmt"

func main() {
	var name string = "omer"
	fmt.Println("name:", name)

	userName := "faruk"
	fmt.Println("user name:", userName)

	var sum int
	fmt.Println("sum:", sum)

	num1, num2 := 5, 55
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)

	sum = num1 + num2
	fmt.Println("sum:", sum)

	var (
		item  = "bottle"
		value = "3"
	)
	fmt.Println("item:", item)
	fmt.Println("value:", value)

	word, _ := "hello", "world"
	fmt.Println("word:", word)
}

package main

import "fmt"

func fizzBuzz(num int) {
	if num%5 == 0 && num%3 == 0 {
		fmt.Println("FizzBuzz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(num)
	}
}

func main() {
	count := 100

	for i := 0; i < count; i++ {
		fizzBuzz(i)
	}
}

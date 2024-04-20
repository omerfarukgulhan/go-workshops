package main

import "fmt"

func double(num int) int {
	return num * 2
}

func add(firstNum, secondNum int) int {
	return firstNum + secondNum
}

func helloWorld() {
	fmt.Println("hello world")
}

func returnDifferent() (int, string) {
	return 5, "string"
}

func main() {
	helloWorld()

	doubleResult := double(12)
	addResult := add(13, 17)

	fmt.Println(doubleResult)
	fmt.Println(addResult)

	num, str := returnDifferent()
	fmt.Println(num)
	fmt.Println(str)
}

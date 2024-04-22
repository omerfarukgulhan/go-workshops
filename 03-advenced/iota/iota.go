package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(num1, mun2 int) int {
	switch op {
	case Add:
		return num1 + mun2
	case Subtract:
		return num1 - mun2
	case Multiply:
		return num1 * mun2
	case Divide:
		return num1 / mun2
	}
	panic("unhandled operation")
}

func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2))

	sub := Operation(Subtract)
	fmt.Println(sub.calculate(10, 3))

	mul := Operation(Multiply)
	fmt.Println(mul.calculate(3, 3))

	div := Operation(Divide)
	fmt.Println(div.calculate(100, 2))
}

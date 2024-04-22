package main

import "fmt"

type Person struct {
	name  string
	age   int
	money int
}

func (person *Person) withDrawMoney(amount int) {
	if person.money <= amount {
		fmt.Println("Not enough money")
	} else {
		person.money -= amount
	}
	fmt.Println("Money:", person.money)
}

func (person *Person) toDepositMoney(amount int) {
	person.money += amount
	fmt.Println("Money:", person.money)
}

func main() {
	person := Person{"omer", 22, 1500}

	person.withDrawMoney(2000)
	person.toDepositMoney(500)
	person.withDrawMoney(2000)
}

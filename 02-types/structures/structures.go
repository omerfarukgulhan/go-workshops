package main

import "fmt"

type Passanger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	firstPassanger Passanger
}

func main() {
	omer := Passanger{"omer", 1, false}
	fmt.Println(omer)

	var (
		ali   = Passanger{"ali", 2, false}
		ahmet = Passanger{"ahmet", 2, false}
	)
	fmt.Println(ali, ahmet)

	omer.Boarded = true
	bus := Bus{omer}

	fmt.Println(bus)
	fmt.Println(bus.firstPassanger.Name)
	fmt.Println(bus.firstPassanger.TicketNumber)

}

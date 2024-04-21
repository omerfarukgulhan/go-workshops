package main

import (
	"fmt"
	"math/rand"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	dice, sides := 2, 6
	rolls := 1

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll:", r, ", die -", d, ":", rolled)
		}
		fmt.Println("Total rolled:", sum)
		switch {
		case sum%2 == 0:
			fmt.Println("Even")
		case sum%2 == 1:
			fmt.Println("Odd")
		}
	}
}

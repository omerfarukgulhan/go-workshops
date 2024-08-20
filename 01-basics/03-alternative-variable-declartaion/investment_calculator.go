package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years, expectedReturnRate := 10000, 10, 4.2

	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Print(futureValue)
}

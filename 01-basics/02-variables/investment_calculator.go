package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 10000
	var expectedReturnRate = 4.2
	var years = 8

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Print(futureValue)
}

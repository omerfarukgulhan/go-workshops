package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue := calcFutureValue(investmentAmount, expectedReturnRate, years)
	futureRealValue := calcFutureRealValue(futureValue, inflationRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func calcFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) float64 {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	return futureValue
}

func calcFutureRealValue(futureValue float64, inflationRate float64, years float64) float64 {
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureRealValue
}

func outputText(text string) {
	fmt.Print(text)
}

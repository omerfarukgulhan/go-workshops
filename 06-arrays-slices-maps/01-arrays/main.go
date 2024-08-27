package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	dynamicPrices := []float64{10.99, 8.99}
	fmt.Println(dynamicPrices[0:1])
	dynamicPrices[1] = 9.99

	dynamicPrices = append(dynamicPrices, 5.99)
	dynamicPrices = dynamicPrices[1:]
	fmt.Println(dynamicPrices)
}

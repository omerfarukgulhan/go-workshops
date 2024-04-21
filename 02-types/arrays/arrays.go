package main

import "fmt"

type ShoppingItem struct {
	name  string
	price int
}

func totalPrice(arr [4]ShoppingItem) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		if arr[i].name != "" {
			fmt.Println(arr[i])
			sum += arr[i].price
		}
	}

	return sum
}

func main() {
	var shoppingList = [4]ShoppingItem{{"apple", 15}, {"banana", 25}, {"orange", 10}}

	fmt.Println(totalPrice(shoppingList))
}

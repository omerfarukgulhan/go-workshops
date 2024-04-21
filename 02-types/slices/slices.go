package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	brand := []string{"audi", "bmw", "mercedes"}
	printSlice("brand 1", brand)

	brand = append(brand, "ford")
	printSlice("brand 2", brand)

	fmt.Println()
	fmt.Println(brand[0], "bought")
	fmt.Println(brand[1], "bought")

	brand = brand[2:]
	printSlice("Remaining brands", brand)
}

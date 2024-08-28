package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	pcNames := make([]string, 2, 5)
	pcNames[0] = "Asus"

	pcNames = append(pcNames, "Hp")
	pcNames = append(pcNames, "Dell")

	fmt.Println(pcNames)

	grades := make(map[string]int, 3)

	grades["AA"] = 90
	grades["AB"] = 80
	grades["BB"] = 70

	fmt.Println(grades)

	for index, value := range pcNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range grades {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}

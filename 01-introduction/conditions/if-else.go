package main

import "fmt"

func avg(a, b, c int) float32 {
	return float32((a + b + c) / 3)
}

func main() {
	quiz1, quiz2, quiz3 := 75, 62, 94

	score := avg(quiz1, quiz2, quiz3)

	fmt.Println("Score:", score)

	if avg(quiz1, quiz2, quiz3) > 70 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	var grade string

	switch {
	case score > 90:
		grade = "AA"
	case score > 80:
		grade = "BB"
	case score > 70:
		grade = "CC"
	case score > 60:
		grade = "DD"
	default:
		grade = "FF"
	}

	fmt.Println("Grade:", grade)
}

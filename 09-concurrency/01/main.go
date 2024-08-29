package main

import (
	"fmt"
	"time"
)

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go test("1", done)
	// dones[1] = make(chan bool)
	go test("2", done)
	// dones[2] = make(chan bool)
	go slowTest("3", done)
	// dones[3] = make(chan bool)
	go test("4", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		// fmt.Println(doneChan)
	}
}

func test(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true
}

func slowTest(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phrase)
	doneChan <- true
	close(doneChan)
}

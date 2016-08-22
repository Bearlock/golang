package main

import (
	"fmt"
)

func incrementingClosure() func() int {
	incrementNumber := 0
	return func() int {
		incrementNumber++
		return incrementNumber
	}
}

func main() {
	firstClosure := incrementingClosure()

	fmt.Println(firstClosure())
	fmt.Println(firstClosure())
	fmt.Println(firstClosure())
	fmt.Println(firstClosure())

	secondClosure := incrementingClosure()

	fmt.Println(secondClosure())
	fmt.Println(secondClosure())
	fmt.Println(secondClosure())
	fmt.Println(secondClosure())
}

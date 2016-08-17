package main

import (
	"fmt"
	"math"
)

const InitialStep = float64(1)
const Delta = .000000001

func newtonsMethod(x float64) float64 {
	
	step := InitialStep
	previousStep := step
	
	for step = newtonsMethodNextStep(x, step); math.Abs(step - previousStep) > Delta; {
		previousStep = step
		step = newtonsMethodNextStep(x, step)
	}
	
	return step
}

func newtonsMethodNextStep(x, step float64) float64 {
	nextStep := (step - ((step * step) - x) / (2 * step))
	return nextStep
}

func main() {
	fmt.Println(newtonsMethod(13))
	fmt.Println(math.Sqrt(13))
}

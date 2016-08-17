package main

import (
	"fmt"
	"math"
)

const InitialStep = float64(1)
const Delta = .000000001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %g", float64(e))
}

func newtonsMethod(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	step := InitialStep
	previousStep := step
	
	for step = newtonsMethodNextStep(x, step); math.Abs(step - previousStep) > Delta; {
		previousStep = step
		step = newtonsMethodNextStep(x, step)
	}
	
	return step, nil
}

func newtonsMethodNextStep(x, step float64) float64 {
	nextStep := (step - ((step * step) - x) / (2 * step))
	return nextStep
}

func main() {
	fmt.Println(newtonsMethod(2))
	fmt.Println(newtonsMethod(-2))
}

package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	dySlice := make([][]uint8, dy)

	for outerIndex := range dySlice {
		dySlice[outerIndex] = make([]uint8, dx)

		for innerIndex := range dySlice[outerIndex] {
			dySlice[outerIndex][innerIndex] = uint8((innerIndex ^ outerIndex) * (innerIndex & outerIndex))
		}

	}
	return dySlice
}

func main() {
	pic.Show(Pic)
}

package main

import (
	"fmt"
	"strconv"
)

type Vertex struct {
	X int
	Y int
}

func vertexStringFormat(v Vertex) string {
	xString := strconv.Itoa(v.X)
	yString := strconv.Itoa(v.Y)
	return "{" + xString + ", " + yString + "}"
}

func main() {
	
	vertexArray := [3]Vertex{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	for i := 0; i < len(vertexArray); i++ {
		fmt.Println(vertexStringFormat(x))
	}
}
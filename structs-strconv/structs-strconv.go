package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func vertexStringFormat(v Vertex) string {
	return fmt.Sprintf("{%d, %d}", v.X, v.Y)
}

func main() {

	vertexArray := [3]Vertex{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	for _, vertex := range vertexArray {
		fmt.Println(vertexStringFormat(vertex))
	}
}
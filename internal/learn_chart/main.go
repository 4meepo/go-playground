package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

func main() {
	data := []float64{200, 390, 800, 1000, 1000, 300}
	graph := asciigraph.Plot(data, asciigraph.Height(8), asciigraph.Width(30), asciigraph.Offset(30))

	fmt.Println(graph)
}

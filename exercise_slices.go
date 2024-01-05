package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	a := make([][]uint8, 0)
	for i := 0; i < dy; i++ {
		a2 := make([]uint8, 0)
		for ii := 0; ii < dx; ii++ {
			a2 = append(a2, uint8((i+11)/2))
		}
		a = append(a, a2)
	}
	return a
}

func main() {
	pic.Show(Pic)
}

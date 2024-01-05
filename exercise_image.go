package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type MyImage struct {
	width  int
	height int
}

func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img MyImage) At(x, y int) color.Color {
	//v := uint8((x + y) / 2)
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := MyImage{300, 300}
	pic.ShowImage(m)
}

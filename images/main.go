package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (Image) At(x, y int) color.Color {
	v := x^y
	return color.RGBA{uint8(v), uint8(v), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
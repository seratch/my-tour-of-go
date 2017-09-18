package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type MyImage struct {
	Width, Height int
	R, G          uint8
}

func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}
func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}
func (img MyImage) At(x, y int) color.Color {
	return color.RGBA{img.R, img.G, 255, 255}
}

func main() {
	m := MyImage{R: 100, G: 100, Width: 100, Height: 200}
	pic.ShowImage(m)
}

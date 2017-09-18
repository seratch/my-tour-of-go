package main

import (
	"fmt"
	"image"
)

func main() {
	/*
		package image
		type Image interface {
		    ColorModel() color.Model
		    Bounds() Rectangle
		    At(x, y int) color.Color
		}
	*/
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	//(0,0)-(100,100)
	//0 0 0 0
}

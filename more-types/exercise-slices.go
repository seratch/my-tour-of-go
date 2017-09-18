package main

// go get golang.org/x/tour/pic
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic = make([][]uint8, dy)
	for y, _ := range pic { // or for y := range pic {
		pic[y] = make([]uint8, dx)
		for x, _ := range pic[y] { // or for x := range pic[y] {
			pic[y][x] = uint8(x ^ y)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}

/*
package main

// go get golang.org/x/tour/pic
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic = make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		var elm = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			elm[x] = uint8(x ^ y)
		}
		pic[y] = elm
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
*/

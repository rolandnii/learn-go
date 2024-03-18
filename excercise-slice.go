package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	image := [][]uint8{}
	for y := 0; y < dy; y++ {
		y_image := []uint8{}
		for x := 0; x < dx; x++ {
			fmt.Printf("The image is %T\n", x^y)
			y_image = append(y_image, uint8(x))
		}
		image = append(image, y_image)
	}

	return image
}

func main() {
	pic.Show(Pic)
}

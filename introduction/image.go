/*
	Author - Anubhav Vats
	Description - Generates beautiful images using shades of blue color when run
				 on Go Playground
	Last Modified - 23/04/21
*/
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		image[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			image[i][j] = uint8(i + j)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
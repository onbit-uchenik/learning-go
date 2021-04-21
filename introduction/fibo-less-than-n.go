/*
	Author - Anubhav Vats
	Description - Get fibonacci number less than any positive integer. Time Complexity - O(n)
	Last Modified - 17/04/21
*/
package main

import (
	"fmt"
)

func getFiboLessThanN(seed int) (int, int) {
	var (
		a = 0
		b = 1
		limit = 500
		iterator = 0
		c = 0
	)
	for iterator < limit {
		c = a + b
		if seed <= c {
			break
		}
		a = b
		b = c
		iterator = iterator + 1
	}
	return a, b
}

func main () {
	fmt.Println(getFiboLessThanN(678))
}
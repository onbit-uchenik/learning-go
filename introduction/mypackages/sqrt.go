/*
	Author - Anubhav Vats
	Description - Calculate the sqrt of any positive number.
	Last Modified - 21/04/21
*/
package mathalgo

import (
	"errors"
	"fmt"
	"math"
	"reflect" // for getting type of any value.
)

func Sqrt(x float64) (error, float64) {
	fmt.Println("checking the type of x", reflect.TypeOf(x)) // prints float64
	if x < 0 {
		return errors.New("expected input to Sqrt to be positive number, found negative number"), 0
	}
	z := x / 2 // assumed root
	limit := 0.00000000000001
	count := 0
	for root := z - ((z*z)-x)/(2*z); math.Abs(root-z) > limit; {
		z = root
		root = z - ((z*z)-x)/(2*z)
		count++
	}
	return nil, z
}

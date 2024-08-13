package main

import (
	"fmt"
)


func Sqrt(x float64) float64{

	// initial value
	z := float64(1)

	// loop
	for i:=0 ; i<10 ; i++{

		//logic
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}

	return z
}

func main(){

	fmt.Println(Sqrt(2))

}
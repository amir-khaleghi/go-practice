package main

/* ■■■■■■■■■■■■■■■■■■■■■■■■ For ■■■■■■■■■■■■■■■■■■■■■■■■ */

// /* Import ----------------------------------------------- */
// import (
// 	"fmt"
// )

// func main() {

// fmt.Println("Hello World")
// sum := 0
// 	for i := 0; i < 10; i++ {

// 		 sum += i
// 	}

// //while loop
//  j :=0
// 	for j< 100{
// 	//run for ever
// 	}

// 	fmt.Println(sum)

// }

/* ■■■■■■■■■■■■■■■■■■■■ If Statement ■■■■■■■■■■■■■■■■■■■■ */
// import (
// 	"fmt"
// 	"math"
// )
// func main (){

// 	fmt.Println(sqrt(9), sqrt(-4))

// /* If Else ---------------------------------------------- */
// a :=2

// if a > 3 {
// 	fmt.Println("Congrats")
// }else{
// 	fmt.Println("You lost the game!")
// }

// }

// func sqrt(x float64) string {
// 	if x < 0 {
//     return sqrt(-x) + "i"
//   }

//   return fmt.Sprintf("%.3f", math.Sqrt(x))
// }

/* ■■■■■■■■■■■■■■■■■■■■ For Continued ■■■■■■■■■■■■■■■■■■■ */
// The init and post statements are optional.

// import (
// 	"fmt"
// )

// func main (){
//  sum := 1
//  for ; sum < 1000; {
// 	sum += sum
// 	fmt.Println(sum)
// }
// }

/* ■■■■■■■■■■■■■■■■■■ For Is Go's While ■■■■■■■■■■■■■■■■■ */

// import "fmt"

// func main(){
// 	sum := 1
// 	for sum < 1000 {
// 		sum += sum
// 		fmt.Println(sum)
// 	}
// }

/* ■■■■■■■■■■■■■■■■■■■■■■■■■ If ■■■■■■■■■■■■■■■■■■■■■■■■■ */

// import (
// 	"fmt"
// 	"math"
// )

// func sqrt(x float64) string {
// 	if x<0 {
// 		return (sqrt(-x)+"i")
// 	}

// 	return fmt.Sprint(math.Sqrt(x))

// }

// func main(){
// 	fmt.Println(sqrt(2), sqrt(-9))
// }

/* ■■■■■■■■■■■■■■■■■■■■■■■ Struct ■■■■■■■■■■■■■■■■■■■■■■■ */

// import "fmt"

// func main (){

// 	/* Define Type ----------------------------------------- */
// 	type Student struct {
// 		Name string
// 		Age  uint
// 		Grade int
// 		ID int
// 	}

// 	/* Use Type -------------------------------------------- */
// 	var amir Student
// 	amir.Name = "Amir"
// 	fmt.Println(amir)

// 	var Eli = Student{
// 		ID: 234,
// 		Name: "Eli",
// 		Age:  20,
// 	}
// 	fmt.Println(Eli)

// 	var Ati = Student{"Ati", 20, 20, 23}

// 	fmt.Println(Ati)

// }

/* ■■■■■■■■■■■■■■■■■■■■■■■ PrintF ■■■■■■■■■■■■■■■■■■■■■■■ */

import "fmt"

func main(){

	type ShoppingCard struct{
		Name string
		Price float64
		Quantity int
		isActive bool
	}


	var card_01 = ShoppingCard{
		Name: "Apple",
		Price: 10.99,
		Quantity: 2,
		isActive : true, // this is a private variable and can't be accessed
	}

	fmt.Printf("The Name of Card is : %s, and The Price: %.2f,and Finally the Quantity: %d\n The process is : %t ", card_01.Name, card_01.Price,card_01.Quantity, card_01.isActive)

}
package main

import "fmt"

func main() {
	// fmt.Println("test fmt")

	/* Types ---------------------------------------------- */
	// var i int = 3
	// var f float64 = 3.43
	// var s stirng = "test"
	// var character rune = "سلام "

	/* Define Variables ----------------------------------- */
	// var firstName string
	// firstName = "Amir"

	// var lastName string = "Khaleghi"

	// var middleName = "Jafar"

	// email := "amir@gmail.com"

	/* Case Sensitivity ----------------------------------- */

	// var name = "a"
	// var Name = "b"

	/* Zero Value ----------------------------------------- */
	// zero value in string is "" in go
	var companyName string

	if companyName == "" {
		fmt.Println("default name")
	} else {
		fmt.Println("company name:", companyName)

	}

	// for rune zero value is 0
	// var r rune

	// for float is 0.0
	// var f float64

	// boolean zero value is false
	var b bool
	if !b {
		fmt.Println("b is false")
	}

	/* ■■■■■■■■■■■■■■■■■■■■■■ Printf ■■■■■■■■■■■■■■■■■■■■■■ */
	templateString := "My name is %s. and my age is %d.\n"
	fmt.Printf(templateString, "Amir", 30)

}

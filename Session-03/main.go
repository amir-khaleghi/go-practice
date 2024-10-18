package main

import (
	"fmt"
	"v1/student"
)

func main() {

	s := student.Student{
		Grades: 32,
	}

	b := student.Test{
		Status: true,
	}

	s.Print()

	b.Testing()

	fmt.Println("Hello World!")
}

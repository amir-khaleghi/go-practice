package student

import "fmt"

/* Print ------------------------------------------------ */
func print() {
	fmt.Println("Student")
}

/* Average Score ---------------------------------------- */
func AvgScore(a1, a2, a3 int) int {
	return (a1 + a2 + a3) / 3
}

package main

import "fmt"

func main() {

	/* For Loop ------------------------------------------- */
	var active = true
	for i := 0; active; i++ {
		fmt.Printf("i=%d\n", i)
		if i == 10 {
			active = false
		}
	}

	/* Make ------------------------------------------------- */

	var sl = make([]int, 5, 10)
	fmt.Printf("sl: %d\n", sl)
	sl[0] = 10
	fmt.Printf("sl: %d\n", sl)

	/* Append --------------------------------------------- */

	var s []int

	fmt.Println("appendme:", s)
	s = append(s, 2, 2, 3)
	fmt.Println("appendme:", s)

	/* Alias ---------------------------------------------- */
	type Amir int
	type ahmad = Amir
	// ahmad is an alias for Amir

}

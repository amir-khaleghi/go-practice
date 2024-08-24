/* ■■■■■■■■■■■■■■■■■■■■■ Os And Log ■■■■■■■■■■■■■■■■■■■■■ */

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	fmt.Println(os.Args)gf
// 	args := os.Args

// 	// error prone

// 	if len(args) < 2 {

// 		// fmt.Println("error - there is no input.")
// 		/* Exit (0) => Successful ----------------------------- */
// 		// os.Exit(1)

// 		/* // No need For Return ----------------------------- */

// 		//instead of 2 above line => use Fatalf

// 		log.Fatalf("There is no in input - len of args : %d", len(args))

// 	}
// 	name := args[1]
// 	fmt.Println("Hello", name)
// }

/* ■■■■■■■■■■■■■■■■■■■■■■■■ Flag ■■■■■■■■■■■■■■■■■■■■■■■■ */

package main

import (
	"flag"
	"fmt"
)

func main() {

	/* Flag ----------------------------------------------- */
	firstName := flag.String("word", "default value", "description")
	age := flag.Int("age", 23, "age of user")

	flag.Parse()
	fmt.Println("name :", *firstName, "age :", *age)

	rFlags := flag.Args()
	fmt.Println("len remaining flags:", len(rFlags))
	fmt.Println("len remaining flags:", rFlags)

	/* Scan ----------------------------------------------- */
	fmt.Printf("please enter your name:")
	var Name, lastName string

	// fmt.Scanln(&Name, &lastName)
	fmt.Println("name:", Name, "LastName: ", lastName)

	/* FormattedString ------------------------------------ */
	fmt.Println("1 2 4 6 8")
	var a1, a2, a3, a4, a5 int

	fmt.Scanf("%d,%d,%d,%d,%d", &a1, &a2, &a3, &a4, &a5)
	fmt.Println(a1, a2, a3, a4, a5)
}

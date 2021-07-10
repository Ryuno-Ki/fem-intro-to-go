// // Uncomment the entire file

package main

import "fmt"

func main() {

	var city string

	switch city {
	case "Des Moines":
		fmt.Println("You live in Iowa")
	// See, an or-structure!
	case "Minneapolis,", "St Paul":
		fmt.Println("You live in Minnesota")
	case "Madison":
		fmt.Println("You live in Wisconsin")
	default:
		fmt.Println("You're not from around here.")
	}

	// 	// ****************************
	var i int

	// Look, no variable here! But conditions in case statements!
	switch {
	case i > 10:
		fmt.Println("Greater than 10")
	case i < 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Is 10")
	}

	// 	// ****************************
	var i int = 9

	switch {
	case i != 10:
		fmt.Println("Does not equal 10")
		// Explicitely mark fall through cases!
		fallthrough
	case i < 10:
		fmt.Println("Less than 10")
	case i > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("Is 10")
	}
}

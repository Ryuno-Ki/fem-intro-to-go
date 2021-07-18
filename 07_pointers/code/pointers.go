package main

import "fmt"

func main() {
	var name string
	// * next to type turns it into a pointer type (memory address)
	var namePointer *string

	fmt.Println("Name:", name)
	fmt.Println("Name *:", namePointer)
}

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
        // Reference memory by /
// 	var namePointer *string = &name
        // Derefence address by *
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

// func changeName(n string) {
// 	n = strings.ToUpper(n)
// }

// func main() {
// 	name := "Elvis"
// 	changeName(name)
// 	fmt.Println(name)
// }

// // ******************************************************

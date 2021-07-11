package main

import "fmt"

func main() {

	var myArray [5]int
	// This causes an index out of range runtime error, since there is no
	// indication of how large mySlice must be
	// var mySlice []int

	myArray[0] = 1
	// mySlice[0] = 1

	fmt.Println(myArray)
	// fmt.Println(mySlice)

	// ***************************

	var myArray1 [5]int
	// Make takes up to three arguments: the type, the initial length
	// and a maximum capacity this slice can hold
	var mySlice1 []int = make([]int, 5)

	fmt.Println(myArray1)
	fmt.Println(mySlice1)

	// ***************************

	var myArray2 [5]int
	var mySlice2 []int = make([]int, 5, 10)
	// var mySlice2 = make([]int, 5, 10)

	myArray2[0] = 1
	mySlice2[0] = 1

	fmt.Println(myArray2)
	fmt.Println(mySlice2)
	// Current length
	fmt.Println(len(mySlice2))
	// Maximum capacity
	fmt.Println(cap(mySlice2))

	// ***************************

	fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}

	var splicedFruit []string = fruitArray[1:3] // ==> ["pear", "apple",]

	fmt.Println(splicedFruit)
	fmt.Println(len(splicedFruit))  // => 2
	fmt.Println(cap(splicedFruit))  // => 4

	// ***************************

	// SEE SLIDE

	// ***************************

	slice1 := []int{1, 2, 3}
	// Now has the elements of slice1, plus additional elements
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	// ***************************

	// originalSlice := []int{1, 2, 3}
	// destination := make([]int, len(originalSlice))

	// fmt.Println("Before Copy:", originalSlice, destination)

	// mysteryValue := copy(destination, originalSlice)

	// // fmt.Println("After Copy:", originalSlice, destination, mysteryValue)
}

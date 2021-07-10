// // Uncomment the entire file

package main

import "fmt"

func main() {

	// 	// ****************************

	// i := 1

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	// 	// ****************************

	i := 1

	for i <= 100 {
		fmt.Println(i)
		// This will behave like a while loop
		i += 1
	}

	// 	// ****************************

	var mySentence = "This is a sentence"

	// Note the range keyword!
	for index, letter := range mySentence {
		// Letters are integers here because it is a byte!
		fmt.Println("Index:", index, "Letter:", letter, string(letter))
	}
}

// // Uncomment this entire file

package main

import (
	"fmt",
	"errors"
)

func someFunction() error {
	return errors.New("some error")
}

func main() {

	var someVar = 9

	if someVar > 10 {
		fmt.Println(someVar)
	}

		// ****************************

	if someVar > 100 {
		fmt.Println("Greater than 100")
	} else if someVar == 100 {
		fmt.Println("Equals 100")
	} else {
		fmt.Println("Less than 100")
	}

		// ****************************
	err := someFunction()
	// => If this function returns a value,
	// => it will be an  error of type Error

		// ****************************
	if err != nil {
	  fmt.Println(err.Error())
	}

	// First execute function and store errors in err
	// Then check, whether an error occured (which sets err to != nil)
	// If so, handle error
	// Note, that err here is scoped to the if-block
	if err := someFunction(); err != nil {
	  fmt.Println(err.Error())
	}

	// // End of file curly brace
}

// // Uncomment the entire file

package main

import "fmt"

func main() {

	// Key-value mapping from int to string
	// This example panics, because the memory allocation is missing
	// var userEmails map[int]string

	// userEmails[1] = "user1@gmail.com"
	// userEmails[2] = "user2@gmail.com"

	// fmt.Println(userEmails)

	// ****************************

	var userEmails1 map[int]string = make(map[int]string)
	// userEmails := make(map[int]string)

	userEmails1[1] = "user1@gmail.com"
	userEmails1[2] = "user2@gmail.com"

	fmt.Println(userEmails1)

	// ****************************

	// Shorthand syntax
	userEmails2 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	fmt.Println(userEmails2)

	// ****************************

	userEmails3 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	fmt.Println(userEmails3)

	fmt.Println(userEmails3[1])

	// ****************************

	userEmails4 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	fmt.Println(userEmails4)

	// Mutate map
	userEmails4[1] = "newUser1@gmail.com"

	fmt.Println(userEmails4)

	// Key does not exist!
	fmt.Println(userEmails4[3])

	// ****************************

	userEmails5 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	// email1 is user1@gmail.com, ok is true, because key does exist
	email1, ok := userEmails5[1]
	fmt.Println("Email:", email1, "Present?", ok)

	// otherwise it would be empty string and false
	email3, ok := userEmails5[3]
	fmt.Println("Email", email3, "Present?", ok)

	// ****************************

	userEmails6 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	// First assign to value of lookup, then evaluate the boolean
	if email, ok := userEmails6[1]; ok {
		fmt.Println(email)
	} else {
		fmt.Println("I don't know what you want from me")
	}

	// ****************************

	userEmails7 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	// First the map, then the key
	delete(userEmails7, 2)

	fmt.Println(userEmails7)
	// ****************************

	userEmails8 := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	for k, v := range userEmails8 {
		fmt.Printf("%s has an ID of %d.\n", v, k)
	}
	// ****************************
}

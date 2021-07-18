package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func updateEmail (u *User) {
  // Structs are handled slightly differently in Go as it allows modification
  // on the memory without explicit refencing and dereferencing
  u.Email = "something@new.com"
}

func main() {
    user := User{ID: 1, FirstName: "Mia", LastName: "Ball", Email: "mia@ball.com"}
    fmt.Println("Original:", user)
    updateEmail(&user)
    fmt.Println("Modified:", user)
}


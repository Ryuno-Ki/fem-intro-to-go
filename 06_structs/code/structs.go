package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
  users          []User
  spaceAvailable bool
}

func describeGroup (g Group) {
  g.spaceAvailable = len(g.users) <= 2
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	users := []User{u, u, u}
	g := Group{ users: users, spaceAvailable: true }
	fmt.Println(g)
}

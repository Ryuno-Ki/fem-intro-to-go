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

// Methods are mainly concerned with state
// So especially use them when updating internal state of an instance of a type
// You can have methods with the same name, but not functions
func (u *User) describe () string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

 func (g *Group) describe () string {
  g.spaceAvailable = len(g.users) <= 2
  return fmt.Sprintf("%t", g.spaceAvailable)
}

func main() {
	user := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	users := []User{user, user, user}
	g := Group{ users: users, spaceAvailable: true }

	// Call mathod
	desc := user.describe()
	fmt.Println(desc)

	fmt.Println(g.describe())
}

package main

import "fmt"

// Add a Describer interface
// By convention, an interface is called by the job it does, normally ending on
// 'er'.
// Instances automatically receive the type if they implement the interface
// (c.f. Duck Typing)
// Interfaces allow for code reduction, by focusing on a common set of method
// signatures
// Therefore, if a struct implements the describe method, its instances will
// be of type Describer. A function can then receive any instance of this
// interface and call the describe method on it
// Instead of having to check for the type
type Describer interface {
  describe() string
}

// There is an empty interface which can be thought of as TypeScript's any:
//interface{}

// You can use it like this:
//var something = map[string]interface{}

// User is a single user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group is a group of Users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

// These two structs have different implementations of the `describe()` method.

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func (g *Group) describe() string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
}

// Create a function that doesn't care what type you pass in as long as the type "satisfies the interface"
// IMPORTANT: No pointer types here!

// Call the describe method on the argument
func Describe (d Describer) string {
  return d.describe()
}

func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 1, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}
	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}

	describeUser := u1.describe()
	describeGroup := g.describe()

	// Pass in reference, because declared on method signature
	describeUserUsingInterface := Describe(&u1)
	describeGroupUsingInterface := Describe(&g)

	fmt.Println(describeUser)
	fmt.Println(describeGroup)
	fmt.Println(describeUserUsingInterface)
	fmt.Println(describeGroupUsingInterface)
}

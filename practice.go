// Every go package needs one and only one package main
// We can pass around variables within the same directory (but not any parent
// directory) to organise our code.
package main

// Access additional packages, here fmt for formatting
import "fmt"

// Entry point of code. Required in every go program
// This one here battles with main.go (same directory). Comment the other one
func main() {
  fmt.Println("Hello World")
}

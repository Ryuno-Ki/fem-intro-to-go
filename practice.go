// Every go package needs one and only one package main
// We can pass around variables within the same directory (but not any parent
// directory) to organise our code.
package main

// Access additional packages, here fmt for formatting
import "fmt"

// Entry point of code. Required in every go program
// This one here battles with main.go (same directory). Comment the other one
func main() {
  // These print statements return number of bytes + write errors
  fmt.Println("Hello World")
  // No newline at the end, unless manually inserted
  fmt.Print("Gibbe")
  fmt.Print("rish\n")
  // Formatted printing. %s is a string verb for strings, %d for integers
  fmt.Printf("Hello, my name is %s. Today is %d July", "André", 10)
  // fmt.Fprint et al. write to file. Returns number of bytes + write errors
  // fmt.Sprint et all store in character buffer and return the string
}

package utils

import "fmt"

func printNum (number int) {
  fmt.Println("Current Number:", number)
}

// Automatically exported, since it starts with an uppercase character

// Adds together multiple numbers
func Add (numbers ...int) int {
  total := 0

  for _, number := range numbers {
    printNum(number)
    total += number
  }

  return total
}

// No func main(), because it is no entry point to our code anymore, but a
// utility library

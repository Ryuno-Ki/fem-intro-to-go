package main

import "fmt"

func main () {
  sentence := "Iterate over me"

  // If index wouldn't be used, Go would yell - unless you declare it using _
  for index, letter := range sentence {
    // index % 2 is not a bool, so either cast it or use a comparison
    if index % 2 == 1 {
      fmt.Println(string(letter))
    }
  }
}

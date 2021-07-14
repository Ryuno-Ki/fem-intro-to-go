package main

import "fmt"

var pets map[string]string = map[string]string{
  "fido": "dog",
  "kitty": "cat",
}

var groceries = []string{"banana", "pear"}

func average (numbers [5]float64) float64 {
  total := 0.0

  for _, number := range numbers {
    total += number
  }

  return total / float64(len(numbers))
}

func is_pet (pet string) bool {
  _, ok := pets[pet]
  return ok
}

func add_groceries (fruits ...string) []string {
  newGroceries := groceries

  for _, fruit := range fruits {
    newGroceries = append(groceries, fruit)
  }

  return newGroceries
}

func main () {
  var scores [5]float64

  fmt.Println(average(scores))

  fmt.Println(pets)
  fmt.Println(is_pet("fido"))
  fmt.Println(is_pet("Frodo"))

  fmt.Println(add_groceries("cherry"))
}

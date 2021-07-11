package main

import "fmt"

// By declaring the variables in the function signature, they don't have to be
// declared within the function, i.e. no var or := needed!
// age1 and age2 resp. ageOfSally and ageOfBob are both of type int
func printAge(age1, age2 int) (ageOfSally, ageOfBob int) {
	ageOfSally = age1
	ageOfBob = age2
	return
}

func printAge1 (age int) int {
  fmt.Println(age)
  return age
}

func printAge2 (age int) (int, int) {
  fmt.Println(age)
  return 0, age
}

// return is still needed, but Go figures out the values from the signature
func printAge3 (age int) (ageOfSally int, ageOfBob int) {
  ageOfBob = 21
  ageOfSally = 16
  return
}

// Accept an unknown number of int values
func printAge4 (ages ...int) {
  // We are not interested in the index, so _ to ignore it
  for _, value := range ages {
    fmt.Println(value)
  }
}

func main() {
	x, y := printAge(10, 21)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(printAge3(8))
	printAge4(16, 21, 30)
}

package main

import "fmt"

var name string = "André"
var surName = "Jaenisch"
var multiple, assignments = "also", "work"

func main () {
  thisIs := "the most common way to assign variables"
  var someString string // Defaults to empty string
  var someNumber int  // Defaults to 0
  var someFloat float64 // Defaults to 0.0
  var someBoolean bool // Defaults to false
  fmt.Println(name, surName, someString, someNumber, someFloat, someBoolean)
  fmt.Println(thisIs)
  fmt.Println(multiple, assignments)
}

package main

import "fmt"

func average (number1, number2, number3 float64) float64 {
  return (number1 + number2 + number3) / 3
}

func main () {
  fmt.Println(average(2, 3, 5))
}

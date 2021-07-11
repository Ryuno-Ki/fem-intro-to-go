package main

import "fmt"

func main () {
  var scores [5]float64

  var scores1 [5]float64
  scores1[0] = 9
  scores1[1] = 1.5
  scores1[2] = 4.5
  scores1[3] = 7
  scores1[4] = 8

  // var scores2 [5]float64 = [5]float64{9, 1.5, 4.5, 7, 8}
  scores3 := [5]float64{9, 1.5, 4.5, 7, 8}
  // scores4 := [...]float64{9, 1.5, 4.5, 7, 8}

  for _, number := range scores3 {
    fmt.Println(number)
  }

  fmt.Println(scores)
}

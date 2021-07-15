package main

// Import by directory structure
// Namespace will be last directory of the path
// Or is it the package declaration at the top of that file?
// Alternatively, add a variable in the line before to act as alias
import (
	"fmt"
	"fem-intro-to-go/05_toolkit/code/utils"
	// math "fem-intro-to-go/05_toolkit/code/utils"
)

func calculateImportantData () int {
  totalValue := utils.Add(1, 2, 3, 4, 5)
  return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateImportantData()
	fmt.Println(total)
}

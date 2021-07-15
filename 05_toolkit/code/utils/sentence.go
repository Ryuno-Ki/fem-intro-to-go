package utils

import "strings"

// MakeExcited transforms a sentence to all caps with an exclamation mark
func MakeExcited (sentence string) string {
  reutrn strings.ToUpper(sentence) + "!"
}

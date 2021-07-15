package utils

import "testing"

func TestAdd (t *testing.T) {
  expected := 9
  actual := utils.Add(2, 3, 4)

  if (expected != actual) {
    t.Errorf("Expected %d but got %d instead", expected, actual)
  }
}

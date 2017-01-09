package test_string

import (
  "fmt"

  "github.com/golang/example/stringutil"
)

func ExampleReverse() {
  fmt.Println(stringutil.Reverse("HOLA"))
  // Output: ALOH!
}
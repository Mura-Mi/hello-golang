package main

import (
  "fmt"
  "math"
)

func add(x int, y int) int {
  return x + y;
}

func main() {
  fmt.Println("My favorite number is", math.Pi);
  fmt.Println("2 + 4 = ", add(2, 4));
}

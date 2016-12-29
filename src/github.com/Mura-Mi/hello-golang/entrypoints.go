package main

import (
  "fmt"
  "math"
)

func add(x, y int) int {
  return x + y;
}

func swap(a, b string) (string, string) {
  return b, a;
}

func main() {
  fmt.Println("My favorite number is", math.Pi);
  fmt.Println("2 + 4 = ", add(2, 4));

  a, b := swap("Hello", "World");
  fmt.Println(a, b);
}

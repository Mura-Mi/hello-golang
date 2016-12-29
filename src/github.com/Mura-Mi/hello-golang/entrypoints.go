package main

import (
  "fmt"
  "math"
)

func swap(a, b string) (string, string) {
  return b, a;
}

func addAndMultiplied(a, b int) (added, multiplied int) {
  added = a + b
  multiplied = a * b
  return
}

func main() {
  fmt.Println("My favorite number is", math.Pi);
  fmt.Println(addAndMultiplied(2, 4));

  a, b := swap("Hello", "World");
  fmt.Println(a, b);

  first, second, third := "Java", "Ruby", "Go"
  fmt.Println(first, second, third);
}

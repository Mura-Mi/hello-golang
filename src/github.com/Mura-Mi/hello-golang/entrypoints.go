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

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }

  return lim;
}

func main() {
  fmt.Println("My favorite number is", math.Pi);
  fmt.Println(addAndMultiplied(2, 4));

  a, b := swap("Hello", "World");
  fmt.Println(a, b);

  first, second, third := "Java", "Ruby", "Go"
  fmt.Println(first, second, third);

  for i := 10; i > 0; i-- {
    defer fmt.Println("defer: ", i);
  }

  fmt.Println(pow(3, 3, 20));
}

package main

import (
  "fmt"
  "math/rand"
  "math"
)

func add(x, y int) int {
  return x + y
}

func main() {
  fmt.Println("My favourite number is", rand.Intn(10))
  fmt.Println(math.Pi)
  fmt.Println(add(42,13))
}

package main

import (
  "fmt"
)

func gen(total int) chan string {
  c := make(chan string)
  go func() {
    for i := 0; i < total; i++ {
      c <- fmt.Sprintf("result %d", i)
    }
    close(c)
  }()
  return c
}
func main() {
  for r := range gen(10) {
    fmt.Println(r)
  }
}

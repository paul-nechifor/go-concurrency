package main

import (
  "fmt"
)

func send(c chan int) {
  for {
    select {
      case c <- 0:
      case c <- 1:
    }
  }
}
func main() {
  c := make(chan int)
  go send(c)
  for i := 0; i < 40; i++ {
    fmt.Printf("%d", <-c)
  }
}

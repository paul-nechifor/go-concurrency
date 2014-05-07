package main

import (
  "fmt"
)

func pass(from, to chan int) {
  to <- 1 + <-from
}
func main() {
  length := 1000000
  var start = make(chan int)
  var a = start
  var b chan int
  for i := 0; i < length; i++ {
    b = make(chan int)
    go pass(a, b)
    a = b
  }
  start <- 0
  fmt.Println(<-b)
}

package main

import (
  "fmt"
)

func worker(id int, c chan string) {
  for {
    fmt.Println("worker", id, <-c)
  }
}
func main() {
  c := make(chan string)
  go worker(0, c)
  go worker(1, c)
  for i := 0; i < 10; i++ {
    c <- fmt.Sprintf("message %d", i)
  }
}

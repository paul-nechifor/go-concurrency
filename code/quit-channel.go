package main

import (
  "fmt"
  "math/rand"
  "time"
)

func fn(msg string) (out chan string, quit chan bool) {
  out, quit = make(chan string), make(chan bool)
  go func() {
    for i := 0; ; i++ {
      time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
      select {
      case out <- fmt.Sprintf("%s %d", msg, i):
      case <-quit: return
      }
    }
  }()
  return
}
func main() {
  result, quit := fn("result")
  for {
    fmt.Println(<-result)
    if rand.Intn(2) == 0 {
      quit <- true
      break
    }
  }
}

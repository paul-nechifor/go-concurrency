package main
import (
  "fmt"
  "time"
)
func hello(s string) {
  for {
    fmt.Println("hello", s)
    time.Sleep(time.Second / 10)
  }
}
func main() {
  go hello("world")
  time.Sleep(time.Second)
}

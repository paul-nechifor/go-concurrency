package main
import (
  "fmt"
)
func other(c chan string) {
  c <- "Hello."
}
func main() {
  c := make(chan string)
  go other(c)
  value := <-c
  fmt.Println(value)
}

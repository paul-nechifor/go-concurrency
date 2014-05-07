package main
import (
  "fmt"
)
func main() {
  c := make(chan string)
  c <- "Hello." // Deadlock here.
  value := <-c

  fmt.Println(value)
}

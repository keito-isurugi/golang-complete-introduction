package main

import (
  "fmt"

  "github.com/keito-isurugi/golang-complete-introduction/module/greeting"
)

func main() {
  message := greeting.Hello("John")
  fmt.Println(message)
}

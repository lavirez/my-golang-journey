package main

import "fmt"

func main() {
  switch x := "bond" {
  case x == "this is the string that":
    fmt.Println("This is the string that this is the string got i printed out")
  }
}

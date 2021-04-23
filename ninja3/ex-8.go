package main

import "fmt"

func main()  {
  switch  {
  case false:
    fmt.Println("This line will not print out!")
  case true:
    fmt.Println("But this one will be printed out")
  }
}

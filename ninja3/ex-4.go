package main

import "fmt"

func main()  {
  i := 1998
  for {
    if i == 2022 {
      break
    }
    fmt.Println(i)
    i++
  }
}

package main

import "fmt"

func main()  {
  for i := 10; i < 101; i++ {
    fmt.Printf("When %v is diveied by 4 the remainder or the modulus is %v\n", i, i%4)
  }
}

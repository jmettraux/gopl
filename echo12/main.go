package main

import (
  "fmt"
  "os"
)

func main() {
  sep := ""
  for i := 0; i < len(os.Args); i++ {
    fmt.Printf("%s%d: %s", sep, i, os.Args[i])
    sep = " "
  }
  fmt.Println("")
}


package main

import "fmt"

func main() {
  var name string
  fmt.Print("Insert your name:")
  fmt.Scanf("%s", &name)
  fmt.Printf("Welcome %s!", name)
}

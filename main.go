package main

import (
  "io/ioutil"
  "fmt"
)


func main() {

  b, err := ioutil.ReadFile("./Seaport.example")
  if err != nil {
    fmt.Print(err)
  }

  images, err := ImageList(b)
  if err != nil {
    fmt.Print(err)
  }

  fmt.Print(images)

}

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

  f, _ := NewFactory("unix://var/run/docker.sock")

  for _, i := range images {
    f.Build(i)
  }
}

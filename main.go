package main

import (
	"fmt"
	"io/ioutil"
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

	for _, image := range images {
		image.Checkout()
	}

	for _, image := range images {
		f.Build(image)
	}
}

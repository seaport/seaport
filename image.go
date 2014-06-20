package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v1"
)

type Image struct {
	Tag      string
	Uri      string
	Context  string
	Path     string
	Build    ImageBuildConfig
	Children []*Image // Images that need to be build before this
}

type ImageBuildConfig struct {
	Before string
	After  string
}

// data => YAML content
func ImageList(data []byte) ([]*Image, error) {
	var config map[string][]*Image

	err := yaml.Unmarshal(data, &config)

	return config["images"], err
}

func (image *Image) Checkout() (string, error) {
	elements := strings.Split(image.Uri, "/")
	name := elements[len(elements)-1]
	fmt.Printf("Checking out %s: %s...\n", name, image.Uri)
	return "", nil
}

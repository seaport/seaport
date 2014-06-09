package main

import (
	"fmt"
	"github.com/libgit2/git2go"
	"gopkg.in/yaml.v1"
	"strings"
)

type Image struct {
	Tag      string
	Uri      string
	Context  string
	Build    ImageBuildConfig
	Children []*Image // Images that need to be build before this
}

type ImageBuildConfig struct {
	Before string
	After  string
	Cache  bool
	Path   string
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
	options := &git.CloneOptions{
		CheckoutBranch: "master",
	}
	fmt.Printf("Checking out %s: %s...\n", name, image.Uri)
	git.Clone(image.Uri, "/tmp/"+name, options)
	return "", nil
}

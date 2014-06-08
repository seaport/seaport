package main

import (
  "gopkg.in/yaml.v1"
)

type Image struct {
  Tag string
  Url string
  Path string
  Build ImageBuildConfig
  Children []*Image // Images that need to be build before this
}

type ImageBuildConfig struct {
  Before string
  After string
  Cache bool
}

// data => YAML content
func ImageList(data []byte) ([]*Image, error) {
  var config map[string][]*Image

  err := yaml.Unmarshal(data, &config)

  return config["images"], err
}

func (i *Image) Checkout() (string, error){
}

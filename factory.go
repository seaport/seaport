package main

import (
  "github.com/fsouza/go-dockerclient"
  "os/exec"
  "log"
  "strings"
  "fmt"
)

type Queue []*Image

type CompletedBuilds []*Image

type Factory struct {
  cache bool
  docker *docker.Client
}

func NewFactory(endpoint string) (*Factory, error) {
  _, err := docker.NewClient(endpoint)
  if err != nil {
    return nil, err
  }

  return &Factory{
    cache: false,
    docker: nil,
  }, nil
}

func (d *Factory) Build(img *Image) bool {
  log.Println("Executing the before command...")
  if err := execute(img.Build.Before); err != nil {
    log.Fatal(err)
  }
  log.Println("Done")

  // docker build

  log.Println("Executing the after command...")
  if err := execute(img.Build.After); err != nil {
    log.Fatal(err)
  }
  log.Println("Done")

  return true
}

func execute(cmd string) error {
  cmd_params := strings.Split(cmd, " ")
  if output, err := exec.Command(cmd_params[0], cmd_params[1:]...).CombinedOutput(); err != nil {
    return fmt.Errorf("Error trying to use git: %s (%s)", err, output)
  }

  return nil
}

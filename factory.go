package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dotcloud/docker/archive"
	"github.com/fsouza/go-dockerclient"
)

type Queue []*Image

type CompletedBuilds []*Image

type Factory struct {
	Cache  bool
	Docker *docker.Client
}

func NewFactory(endpoint string) (*Factory, error) {
	e, err := docker.NewClient(endpoint)
	if err != nil {
		return nil, err
	}

	return &Factory{
		Cache:  false,
		Docker: e,
	}, nil
}

func (f *Factory) Build(img *Image) bool {
	p, err := filepath.Abs(img.Path)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if _, err := os.Stat(filepath.Join(p, "Dockerfile")); err != nil {
		log.Fatalf("There is no Docker file in %v or is not accessible", p)
		log.Fatal(err)
		return false
	}

	if err := os.Chdir(p); err != nil {
		log.Fatal(err)
		return false
	}

	log.Println("Executing the before command...")
	if err := execute(img.Build.Before); err != nil {
		log.Fatal(err)
		return false
	}
	log.Println("Done")

	log.Println("Preparing the context...")
	r, err := archive.Tar(img.Path, archive.Uncompressed)
	if err != nil {
		log.Fatal(err)
		return false
	}

	out := bytes.NewBufferString("")
	opts := docker.BuildImageOptions{
		Name:           img.Tag,
		NoCache:        !f.Cache,
		RmTmpContainer: true,
		InputStream:    r,
		OutputStream:   out,
	}

	log.Println("Building image...")
	if err := f.Docker.BuildImage(opts); err != nil {
		log.Fatal(err)
		return false
	}
	log.Println(out)
	log.Println("Done")

	log.Println("Executing the after command...")
	if err := execute(img.Build.After); err != nil {
		log.Fatal(err)
		return false
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

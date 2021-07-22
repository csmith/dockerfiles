package main

import (
	"log"

	"github.com/csmith/dockerfiles"
	"github.com/google/go-containerregistry/pkg/crane"
)

const (
	baseImage = "reg.c5h.io/alpine"
)

func main() {
	digest, err := crane.Digest(baseImage)
	if err != nil {
		log.Fatalf("Couldn't determine base image version: %v", err)
	}

	if err := dockerfiles.RenderTemplate("base/Dockerfile", dockerfiles.BaseImage{
		Ref:    baseImage,
		Digest: digest,
	}); err != nil {
		log.Fatalf("Unable to generate template: %v", err)
	}
}

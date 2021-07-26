package main

import (
	"log"

	"github.com/csmith/dockerfiles"
)

const (
	baseImage = "reg.c5h.io/alpine"
)

func main() {
	digest, err := dockerfiles.LatestDigest(baseImage)
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

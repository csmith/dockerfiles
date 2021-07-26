package main

import (
	"log"

	"github.com/csmith/dockerfiles"
)

const (
	golangImage = "reg.c5h.io/golang"
	baseImage   = "reg.c5h.io/base"
)

type templateArgs struct {
	Golang dockerfiles.BaseImage
	Base   dockerfiles.BaseImage
}

func main() {
	// There's no point checking for actual updates yet, as the next release will break the build process
	// because the project migrated from docker/distribution to distribution/distribution.

	baseDigest, err := dockerfiles.LatestDigest(baseImage)
	if err != nil {
		log.Fatalf("Couldn't determine base image version: %v", err)
	}

	golangDigest, err := dockerfiles.LatestDigest(golangImage)
	if err != nil {
		log.Fatalf("Couldn't determine base image version: %v", err)
	}

	if err := dockerfiles.RenderTemplate("distribution/Dockerfile", templateArgs{
		Base: dockerfiles.BaseImage{
			Ref:    baseImage,
			Digest: baseDigest,
		},
		Golang: dockerfiles.BaseImage{
			Ref:    golangImage,
			Digest: golangDigest,
		},
	}); err != nil {
		log.Fatalf("Unable to generate template: %v", err)
	}
}

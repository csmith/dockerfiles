package main

import (
	"log"

	"github.com/csmith/dockerfiles"
)

const (
	baseImage = "reg.c5h.io/alpine"
)

type templateArgs struct {
	Base     dockerfiles.BaseImage
	Packages map[string]string
}

func main() {
	digest, err := dockerfiles.LatestDigest(baseImage)
	if err != nil {
		log.Fatalf("Couldn't determine base image version: %v", err)
	}

	packages, err := dockerfiles.LatestAlpinePackages("ca-certificates", "musl", "tzdata", "rsync")
	if err != nil {
		log.Fatalf("Unable to retrieve latest packages: %v", err)
	}

	if err := dockerfiles.RenderTemplate("base/Dockerfile", templateArgs{
		Base: dockerfiles.BaseImage{
			Ref:    baseImage,
			Digest: digest,
		},
		Packages: packages,
	}); err != nil {
		log.Fatalf("Unable to generate template: %v", err)
	}
}

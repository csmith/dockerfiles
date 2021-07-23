package main

import (
	"log"

	"github.com/csmith/dockerfiles"
	"github.com/google/go-containerregistry/pkg/crane"
)

const (
	baseImage   = "reg.c5h.io/alpine"
	golangImage = "reg.c5h.io/golang"
	repo        = "hashicorp/vault"
)

type templateArgs struct {
	Build dockerfiles.BaseImage
	Run   dockerfiles.BaseImage
	Tag   string
}

func main() {
	baseDigest, err := crane.Digest(baseImage)
	if err != nil {
		log.Fatalf("Couldn't determine base image version: %v", err)
	}

	golangDigest, err := crane.Digest(golangImage)
	if err != nil {
		log.Fatalf("Couldn't determine golang image version: %v", err)
	}

	tag, err := dockerfiles.LatestGitHubTag(repo)
	if err != nil {
		log.Fatalf("Couldn't determine latest tag: %v", err)
	}

	if err := dockerfiles.RenderTemplate("vault/Dockerfile", templateArgs{
		Build: dockerfiles.BaseImage{
			Ref:    golangImage,
			Digest: golangDigest,
		},
		Run: dockerfiles.BaseImage{
			Ref:    baseImage,
			Digest: baseDigest,
		},
		Tag: tag,
	}); err != nil {
		log.Fatalf("Unable to generate template: %v", err)
	}
}

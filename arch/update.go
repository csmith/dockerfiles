package main

import (
	"log"

	"github.com/csmith/dockerfiles"
)

type templateArgs struct {
	Packages map[string]string
}

func main() {
	packages, err := dockerfiles.LatestArchPackages("ca-certificates", "filesystem", "glibc", "gzip", "licenses", "pacman", "sed")
	if err != nil {
		log.Fatalf("Unable to retrieve latest packages: %v", err)
	}

	if err := dockerfiles.RenderTemplate("arch/Dockerfile", templateArgs{
		Packages: packages,
	}); err != nil {
		log.Fatalf("Unable to generate template: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/csmith/dockerfiles"
	"github.com/hashicorp/go-version"
)

const (
	baseImage    = "reg.c5h.io/alpine"
	releaseIndex = "https://ftp.postgresql.org/pub/source/"
	downloadUrl  = "https://ftp.postgresql.org/pub/source/v%[1]s/postgresql-%[1]s.tar.bz2"
	checksumUrl  = "https://ftp.postgresql.org/pub/source/v%[1]s/postgresql-%[1]s.tar.bz2.sha256"
)

type templateArgs struct {
	Image   dockerfiles.BaseImage
	Archive dockerfiles.UrlAndSum
}

func main() {
	digest, err := dockerfiles.LatestDigest(baseImage)
	if err != nil {
		log.Fatalf("Couldn't determine base image version: %v", err)
	}

	versions, err := dockerfiles.FindInHtml(releaseIndex, `a[href*="v13."]`)
	if err != nil {
		log.Fatalf("Couldn't find releases: %v", err)
	}

	best := version.Must(version.NewVersion("0.0.0"))
	found := ""
	for i := range versions {
		v := version.Must(version.NewVersion(strings.TrimSuffix(versions[i], "/")))
		if (best == nil || v.GreaterThanOrEqual(best)) && v.Prerelease() == "" {
			best = v
			found = strings.TrimPrefix(strings.TrimSuffix(versions[i], "/"), "v")
		}
	}

	if found == "" {
		log.Fatalf("Couldn't find candidate version from releases: %v", versions)
	}

	hash, err := dockerfiles.DownloadHash(fmt.Sprintf(checksumUrl, found))

	if err := dockerfiles.RenderTemplate("postgres-13/Dockerfile", templateArgs{
		Image: dockerfiles.BaseImage{
			Ref:    baseImage,
			Digest: digest,
		},
		Archive: dockerfiles.UrlAndSum{
			Url: fmt.Sprintf(downloadUrl, found),
			Sum: hash,
		},
	}); err != nil {
		log.Fatalf("Unable to generate template: %v", err)
	}
}

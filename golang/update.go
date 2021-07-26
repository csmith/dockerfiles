package main

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/csmith/dockerfiles"
	"github.com/hashicorp/go-version"
)

const (
	baseImage = "reg.c5h.io/alpine"
	baseUrl   = "https://golang.org/dl/"
	jsonUrl   = baseUrl + "?mode=json"
	fileKind  = "source"
)

type templateArgs struct {
	Image dockerfiles.BaseImage
	File  dockerfiles.UrlAndSum
}

func main() {
	url, sum, err := latestFile()
	if err != nil {
		log.Fatalf("Unable to retrieve latest file: %v", err)
	}

	digest, err := dockerfiles.LatestDigest(baseImage)
	if err != nil {
		log.Fatalf("Couldn't determine base image version: %v", err)
	}

	if err := dockerfiles.RenderTemplate("golang/Dockerfile", templateArgs{
		Image: dockerfiles.BaseImage{
			Ref:    baseImage,
			Digest: digest,
		},
		File: dockerfiles.UrlAndSum{
			Url: url,
			Sum: sum,
		},
	}); err != nil {
		log.Fatalf("Unable to generate template: %v", err)
	}
}

func latestFile() (string, string, error) {
	var releases []struct {
		Version string `json:"version"`
		Files   []struct {
			Filename string `json:"filename"`
			Checksum string `json:"sha256"`
			Kind     string `json:"kind"`
		} `yaml:"files"`
	}

	if err := dockerfiles.DownloadJson(jsonUrl, &releases); err != nil {
		return "", "", err
	}

	best := version.Must(version.NewVersion("0.0.0"))
	for i := range releases {
		r := releases[i]
		v := version.Must(version.NewVersion(strings.TrimPrefix(r.Version, "go")))
		if v.GreaterThanOrEqual(best) {
			best = v

			for j := range r.Files {
				if r.Files[j].Kind == fileKind {
					return path.Join(baseUrl, r.Files[j].Filename), r.Files[j].Checksum, nil
				}
			}
		}
	}

	return "", "", fmt.Errorf("no release found matching '%s'", fileKind)
}

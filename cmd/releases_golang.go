package main

import (
	"log"
	"path"
	"strings"

	"github.com/csmith/dockerfiles"
	"github.com/hashicorp/go-version"
)

func init() {
	AddRelease("golang", func() (latest string, url string, checksum string) {
		const (
			golangBaseUrl  = "https://golang.org/dl/"
			golangJsonUrl  = golangBaseUrl + "?mode=json"
			golangFileKind = "source"
		)

		var releases []struct {
			Version string `json:"version"`
			Files   []struct {
				Filename string `json:"filename"`
				Checksum string `json:"sha256"`
				Kind     string `json:"kind"`
			} `yaml:"files"`
		}

		if err := dockerfiles.DownloadJson(golangJsonUrl, &releases); err != nil {
			log.Fatalf("Unable to download golang release information: %v", err)
		}

		best := version.Must(version.NewVersion("0.0.0"))
		for i := range releases {
			r := releases[i]
			v := version.Must(version.NewVersion(strings.TrimPrefix(r.Version, "go")))
			if v.GreaterThanOrEqual(best) {
				best = v

				for j := range r.Files {
					if r.Files[j].Kind == golangFileKind {
						latest = r.Version
						url = path.Join(golangBaseUrl, r.Files[j].Filename)
						checksum = r.Files[j].Checksum
					}
				}
			}
		}

		if latest == "" {
			log.Fatalf("No golang release found matching '%s'", golangFileKind)
		}
		return
	})
}

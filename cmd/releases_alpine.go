package main

import (
	"log"

	"github.com/csmith/dockerfiles"
)

func init() {
	AddRelease("alpine", func() (latest string, url string, checksum string) {
		const (
			alpineBaseUrl      = "https://mirrors.melbourne.co.uk/alpine/latest-stable/releases/x86_64/"
			alpineReleaseIndex = alpineBaseUrl + "latest-releases.yaml"
			alpineReleaseTitle = "Mini root filesystem"
		)

		var releases []struct {
			Title    string `yaml:"title"`
			File     string `yaml:"file"`
			Checksum string `yaml:"sha256"`
			Version  string `yaml:"version"`
		}

		if err := dockerfiles.DownloadYaml(alpineReleaseIndex, &releases); err != nil {
			log.Fatalf("Unable to download Alpine release information: %v", err)
		}

		for i := range releases {
			if releases[i].Title == alpineReleaseTitle {
				return releases[i].Version, alpineBaseUrl + releases[i].File, releases[i].Checksum
			}
		}

		log.Fatalf("No Alpine release found matching '%s'", alpineReleaseTitle)
		return
	})
}

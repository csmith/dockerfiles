package main

import (
	"fmt"
	"log"

	"github.com/csmith/dockerfiles"
)

const (
	baseUrl   = "https://dl-cdn.alpinelinux.org/alpine/latest-stable/releases/x86_64/"
	yamlUrl   = baseUrl + "latest-releases.yaml"
	distTitle = "Mini root filesystem"
)

func main() {
	url, sum, err := latestFile()
	if err != nil {
		log.Fatalf("Unable to retrieve latest file: %v", err)
	}

	if err := dockerfiles.RenderTemplate("alpine/Dockerfile", dockerfiles.UrlAndSum{
		Url: url,
		Sum: sum,
	}); err != nil {
		log.Fatalf("Unable to generate template: %v", err)
	}
}

func latestFile() (string, string, error) {
	var releases []struct {
		Title    string `yaml:"title"`
		File     string `yaml:"file"`
		Checksum string `yaml:"sha256"`
	}

	if err := dockerfiles.DownloadYaml(yamlUrl, &releases); err != nil {
		return "", "", err
	}

	for i := range releases {
		if releases[i].Title == distTitle {
			return baseUrl + releases[i].File, releases[i].Checksum, nil
		}
	}

	return "", "", fmt.Errorf("no release found matching '%s'", distTitle)
}

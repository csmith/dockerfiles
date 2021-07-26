package dockerfiles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/hashicorp/go-version"
	"gopkg.in/yaml.v2"
)

// LatestDigest finds the latest digest for the given docker image reference.
func LatestDigest(ref string) (string, error) {
	var authOpt crane.Option

	if *registryUser == "" || *registryPass == "" {
		authOpt = crane.WithAuthFromKeychain(authn.DefaultKeychain)
	} else {
		authOpt = crane.WithAuth(&authn.Basic{
			Username: *registryUser,
			Password: *registryPass,
		})
	}

	return crane.Digest(ref, authOpt)
}

// DownloadYaml requests the given url and then attempts to unmarshal the body as YAML into the provided struct.
func DownloadYaml(url string, i interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	return yaml.NewDecoder(r.Body).Decode(i)
}

// DownloadJson requests the given url and then attempts to unmarshal the body as JSON into the provided struct.
func DownloadJson(url string, i interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(i)
}

// LatestGitHubTag uses the GitHub API to find the tag for the latest stable release.
func LatestGitHubTag(repo string) (string, error) {
	var releases []struct {
		Name string `json:"name"`
	}

	if err := DownloadJson(fmt.Sprintf("https://api.github.com/repos/%s/tags", repo), &releases); err != nil {
		return "", err
	}

	best := version.Must(version.NewVersion("0.0.0"))
	bestTag := ""
	for i := range releases {
		v, err := version.NewVersion(releases[i].Name)
		if err == nil && v.GreaterThanOrEqual(best) && v.Prerelease() == "" {
			best = v
			bestTag = releases[i].Name
		}
	}

	if bestTag == "" {
		return "", fmt.Errorf("no stable semver tags found")
	}
	return bestTag, nil
}

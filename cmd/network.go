package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	return crane.Digest(fmt.Sprintf("%s/%s", *registry, ref), authOpt)
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

// DownloadHash downloads the given URL and parses the first hash out of it, assuming it's formatted in line with the
// output of sha256sum. Hashes are assumed to be hexadecimal and an error will be returned if this is not the case.
func DownloadHash(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	hash := strings.ToLower(strings.SplitN(string(b), " ", 2)[0])
	for i := range hash {
		if (hash[i] < 'a' || hash[i] > 'f') && (hash[i] < '0' || hash[i] > '9') {
			return "", fmt.Errorf("invalid has found at address: %s", hash)
		}
	}
	return hash, nil
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

// FindInHtml downloads the HTML page at the given URL and runs the specified CSS selector over it to find nodes.
// The textual content of those nodes is returned.
func FindInHtml(url string, selector string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var results []string
	doc.Find(selector).Each(func(i int, selection *goquery.Selection) {
		results = append(results, selection.Text())
	})
	return results, nil
}

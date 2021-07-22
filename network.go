package dockerfiles

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

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

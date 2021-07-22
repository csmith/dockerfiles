package dockerfiles

import (
	"fmt"
	"os"
	"text/template"
)

// UrlAndSum contains a URL and a checksum for the file at that URL, for use in templates.
type UrlAndSum struct {
	Url string
	Sum string
}

// BaseImage contains a docker image ref and the digest that should be used.
type BaseImage struct {
	Ref    string
	Digest string
}

// RenderTemplate creates a file at the given target by parsing the `${target}.gotpl` template and rendering it with
// the given args.
func RenderTemplate(target string, args interface{}) error {
	tpl, err := template.ParseFiles(fmt.Sprintf("%s.gotpl", target))
	if err != nil {
		return err
	}

	o, err := os.Create(target)
	if err != nil {
		return err
	}
	defer o.Close()

	return tpl.Execute(o, args)
}

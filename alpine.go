package dockerfiles

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apkIndexUrl = "https://mirrors.melbourne.co.uk/alpine/latest-stable/main/x86_64/APKINDEX.tar.gz"

// LatestAlpinePackages returns a map of packages to their latest version. The result will include all of the provided
// package names, plus all of their direct and transitive dependencies.
func LatestAlpinePackages(names ...string) (map[string]string, error) {
	packages, err := apkPackageInfos()
	if err != nil {
		return nil, err
	}

	res := make(map[string]string)
	queue := append([]string{}, names...)

	for len(queue) > 0 {
		if _, ok := res[queue[0]]; ok {
			// We've already got a resolution for this package, skip it.
			queue = queue[1:]
			continue
		}

		p, ok := packages[queue[0]]
		if !ok {
			return nil, fmt.Errorf("package required but not found: %s", queue[0])
		}

		queue = append(queue[1:], p.Dependencies...)
		res[p.Name] = p.Version
	}

	return res, nil
}

var apkPackageCache map[string]*packageInfo

// apkPackageInfos returns a map of all apk pacman packages and their latest info.
func apkPackageInfos() (map[string]*packageInfo, error) {
	if apkPackageCache != nil {
		return apkPackageCache, nil
	}

	apkPackageCache = make(map[string]*packageInfo)
	res, err := http.Get(apkIndexUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	info, err := readApkIndex(res.Body)
	if err != nil {
		return nil, err
	}
	for k := range info {
		apkPackageCache[k] = info[k]
	}

	return apkPackageCache, nil
}

// readApkIndex reads a .tar.gz archive containing an APKINDEX file, returning the packages within.
func readApkIndex(reader io.Reader) (map[string]*packageInfo, error) {
	gz, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	res := make(map[string]*packageInfo)
	tr := tar.NewReader(gz)
	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		if header.Typeflag == tar.TypeReg && header.Name == "APKINDEX" {
			return readApkIndexContent(tr), nil
		}
	}
	return res, nil
}

// readApkIndexContent reads an APKINDEX file, parsing out the contained packages.
func readApkIndexContent(reader io.Reader) map[string]*packageInfo {
	res := make(map[string]*packageInfo)
	scanner := bufio.NewScanner(reader)

	current := &packageInfo{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			res[current.Name] = current

			for i := range current.Provides {
				// Don't overwrite real packages with provides info
				if _, ok := res[current.Provides[i]]; !ok {
					res[current.Provides[i]] = current
				}
			}

			current = &packageInfo{}
		} else if strings.HasPrefix(line, "P:") {
			current.Name = strings.TrimPrefix(line, "P:")
		} else if strings.HasPrefix(line, "D:") {
			d := strings.Fields(strings.TrimPrefix(line, "D:"))
			for i := range d {
				current.Dependencies = append(current.Dependencies, stripVersion(d[i]))
			}
		} else if strings.HasPrefix(line, "p:") {
			p := strings.Fields(strings.TrimPrefix(line, "p:"))
			for i := range p {
				current.Provides = append(current.Provides, stripVersion(p[i]))
			}
		} else if strings.HasPrefix(line, "V:") {
			current.Version = strings.TrimPrefix(line, "V:")
		}
	}

	return res
}


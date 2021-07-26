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

const pacmanDatabaseUrl = "https://mirrors.melbourne.co.uk/archlinux/%[1]s/os/x86_64/%[1]s.db"

// LatestArchPackages returns a map of packages to their latest version. The result will include all of the provided
// package names, plus all of their direct and transitive dependencies.
func LatestArchPackages(names ...string) (map[string]string, error) {
	packages, err := packageInfos()
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

var packageCache map[string]*packageInfo

// packageInfos returns a map of all known pacman packages and their latest info.
func packageInfos() (map[string]*packageInfo, error) {
	if packageCache != nil {
		return packageCache, nil
	}

	packageCache = make(map[string]*packageInfo)
	for _, repo := range []string{"community", "extra", "core"} {
		err := func() error {
			res, err := http.Get(fmt.Sprintf(pacmanDatabaseUrl, repo))
			if err != nil {
				return err
			}
			defer res.Body.Close()
			info, err := readDatabase(res.Body)
			if err != nil {
				return err
			}
			for k := range info {
				packageCache[k] = info[k]
			}
			return nil
		}()
		if err != nil {
			return nil, err
		}
	}

	return packageCache, nil
}

// readDatabase reads all information from a pacman package database
func readDatabase(reader io.Reader) (map[string]*packageInfo, error) {
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

		if header.Typeflag == tar.TypeReg {
			info, err := readPackageInfo(tr)
			if err != nil {
				return nil, fmt.Errorf("error reading %s: %v", header.Name, err)
			}
			res[info.Name] = info
			for i := range info.Provides {
				res[info.Provides[i]] = info
			}
		}
	}
	return res, nil
}

// packageInfo describes a package available in a pacman repository.
type packageInfo struct {
	Name         string
	Version      string
	Dependencies []string
	Provides     []string
}

// readPackageInfo reads a `desc` file from within the pacman package database, parsing out the relevant information.
func readPackageInfo(reader io.Reader) (*packageInfo, error) {
	res := &packageInfo{}
	scanner := bufio.NewScanner(reader)
	section := ""
	n := 0

	for scanner.Scan() {
		line := scanner.Text()
		n++
		if section == "" {
			if strings.HasPrefix(line, "%") {
				section = line
			} else {
				return nil, fmt.Errorf("line %d: expected section name, got '%s'", n, line)
			}
		} else if line == "" {
			section = ""
		} else if section == "%NAME%" {
			res.Name = line
		} else if section == "%VERSION%" {
			res.Version = line
		} else if section == "%PROVIDES%" {
			res.Provides = append(res.Provides, stripVersion(line))
		} else if section == "%DEPENDS%" {
			res.Dependencies = append(res.Dependencies, stripVersion(line))
		}
	}

	return res, nil
}

// stripVersion removes version qualifiers from a package name such as `foo>=1.2`.
func stripVersion(name string) string {
	i := strings.IndexAny(name, ">=<")
	if i > -1 {
		return name[0:i]
	}
	return name
}

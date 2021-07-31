package main

import (
	"fmt"
	"sync"
)

func AddRelease(name string, fun func() (version string, url string, checksum string)) {
	var once sync.Once
	var url, checksum string

	init := func() {
		once.Do(func() {
			var version string
			version, url, checksum = fun()
			materials[name] = version
		})
	}

	funcs[fmt.Sprintf("%s_url", name)] = func() string {
		init()
		return url
	}
	funcs[fmt.Sprintf("%s_checksum", name)] = func() string {
		init()
		return checksum
	}
}

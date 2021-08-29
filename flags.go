package dockerfiles

import "flag"

var (
	registry     = flag.String("registry", "reg.c5h.io", "Registry to use for pushes and pulls")
	registryUser = flag.String("registry-user", "", "Username to use when talking to docker registries")
	registryPass = flag.String("registry-pass", "", "Password to use when talking to docker registries")
)

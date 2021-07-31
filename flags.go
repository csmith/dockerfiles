package dockerfiles

import "flag"

var (
	registryUser = flag.String("registry-user", "", "Username to use when talking to docker registries")
	registryPass = flag.String("registry-pass", "", "Password to use when talking to docker registries")
)

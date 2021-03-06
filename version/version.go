package version

import (
	"fmt"
)

const Major = "0"
const Minor = "1"
const Patch = "2"
const Label = "beta"

var (
	// Version is the current version of Echoin
	Version string

	// GitCommit is set with --ldflags "-X main.gitCommit=$(git rev-parse --short HEAD)"
	GitCommit string
)

func init() {
	Version = fmt.Sprintf("%s.%s.%s", Major, Minor, Patch)

	if Label != "" {
		Version += "-" + Label
	}

	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}

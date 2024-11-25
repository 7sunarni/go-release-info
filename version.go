package version

import "log"

var (
	GitMajor string
	GitMinor string

	GitVersion   = "v0.0.0-master+$Format:%H$"
	GitCommit    = "$Format:%H$"
	GitTreeState = ""

	BuildDate = "1970-01-01T00:00:00Z"
)

func init() {
	log.Printf("GitVersion: %s GitCommit: %s", GitVersion, GitCommit)
}

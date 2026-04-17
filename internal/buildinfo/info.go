package buildinfo

import "fmt"

var (
	Version   = "dev"
	Commit    = "unknown"
	BuildTime = "unknown"
)

func String() string {
	return fmt.Sprintf("Version: \t%s\nCommit: \t%s\nBuildTime: \t%s", Version, Commit, BuildTime)
}

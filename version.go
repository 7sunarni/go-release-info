package version

import (
	"fmt"
	"os"
)

var (
	GitBranch = "default"
	GitCommit = "default"
	BuildDate = "default"
)

func init() {
	if GitBranch == "default" && GitCommit == "default" && BuildDate == "default" {
		fmt.Fprint(os.Stdout, `You used github.com/7sunarni/go-release-info to generate release information but not assigned values. 
You can add below command to generate your release information.
GitBranch=$(git branch --show-current)
GitCommit=$(git rev-parse "HEAD^{commit}")
BuildDate=$(date +%Y-%m-\%dT%H-%M-%S)
ARGS="-X 'github.com/7sunarni/go-release-info.GitBranch="$GitBranch"'"
ARGS=$ARGS" -X 'github.com/7sunarni/go-release-info.GitCommit="$GitCommit"'" 
ARGS=$ARGS" -X 'github.com/7sunarni/go-release-info.BuildDate="$BuildDate"'"`)
	} else {
		fmt.Fprintf(os.Stdout, `Below is release version information generated by the github.com/7sunarni/go-release-info.
Your project GitBranch is: %s
GitCommit is: %s
Built in %s.`, GitBranch, GitCommit, BuildDate)
	}

}

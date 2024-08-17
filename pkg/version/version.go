package version

import (
	"os/exec"
    "strings"

    "k8s.io/klog/v2"
)

var (
	GitCommit string

	Version string
)

func GetGitCommit() string {
	if GitCommit != "" {
		return GitCommit
	}

	cmd := exec.Command("git", "rev-parse", "--verify", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		klog.Errorf("failed to get git commit: %s", err.Error())
		return ""
	}

	return strings.TrimSpace(string(output))
}

func GetVersion() string {
	return "v0.0.1"
}
package extconfig

import (
	"os"
	"strings"
)

// Set via ldflags
var (
	Version   = "99.99.99"
	BuildHash = ""
)

const ReleaseBuildHash = "release"

func CI() bool {
	return strings.EqualFold(os.Getenv("CI"), "true")
}

func FullVersion() string {
	if CI() {
		Version = "99.99.99"
		BuildHash = "test"
	}
	version := strings.Builder{}
	version.WriteString(Version)
	if BuildHash != "" && BuildHash != ReleaseBuildHash {
		version.WriteRune('+')
		if len(BuildHash) >= 8 {
			version.WriteString(BuildHash[:8])
		} else {
			version.WriteString(BuildHash)
		}
	}
	return version.String()
}

func init() {
	if BuildHash == "" {
		BuildHash = "dev"
	}
}

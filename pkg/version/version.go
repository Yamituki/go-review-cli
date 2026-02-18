package version

import "fmt"

var (
	Version   = "1.0.0"
	GitCommit = "dev"
	BuildDate = "unknown"
)

// GetVersion はバージョン情報、Gitコミットハッシュ、ビルド日時を返します。
func GetVersion() string {
	return Version
}

// GetFullVersion は完全なバージョン情報を返します。
func GetFullVersion() string {
	return fmt.Sprintf("%s (commit: %s, built: %s)", Version, GitCommit, BuildDate)
}

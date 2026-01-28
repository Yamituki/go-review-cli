package version

import "fmt"

var (
	Version   string = "0.1.0"
	GitCommit string = "dev"
	BuildDate string = "unknown"
)

// GetVersion はバージョン情報、Gitコミットハッシュ、ビルド日時を返します。
func GetVersion() string {
	return Version
}

// GetFullVersion は完全なバージョン情報を返します。
func GetFullVersion() string {
	return fmt.Sprintf("%s (commit: %s, built: %s)", Version, GitCommit, BuildDate)
}

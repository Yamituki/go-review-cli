package version

/**
 * fmt パッケージはフォーマットされたI/Oを提供します。
 */
import "fmt"

/**
 * これらの変数はビルド時に設定されます。
 * Version はビルド時に設定されるアプリケーションのバージョン情報を保持します。
 * GitCommit はビルド時に設定されるGitコミットハッシュを保持します。
 * BuildDate はビルド時に設定されるビルド日時を保持します。
 */
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

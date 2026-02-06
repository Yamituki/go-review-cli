package commands

import (
	"fmt"

	"github.com/Yamituki/go-review-cli/pkg/version"
	"github.com/spf13/cobra"
)

// InitVersionCommand バージョン情報を表示するコマンドを初期化する
func InitVersionCommand(root *cobra.Command) {
	// バージョンコマンドを作成
	versionCommand := &cobra.Command{
		Use:   "version",
		Short: "バージョン情報を表示",
		Long:  "go-review-cliのバージョン、Gitコミット、ビルド日時を表示します",
		Run:   runVersion,
	}

	// ルートコマンドにバージョンコマンドを追加
	root.AddCommand(versionCommand)
}

// runVersion バージョン情報を表示する関数
func runVersion(cmd *cobra.Command, args []string) {
	// バージョン情報を取得　(v0.1.0 (commit: dev, built: unknown))
	// versionInfo := version.GetFullVersion()

	// バージョン情報を出力
	cmd.Println(fmt.Sprintf("go-review-cli version: %s", version.Version))
	cmd.Println(fmt.Sprintf("commit: %s", version.GitCommit))
	cmd.Println(fmt.Sprintf("built: %s", version.BuildDate))
}

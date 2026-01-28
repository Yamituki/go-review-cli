package commands

import (
	"fmt"

	"github.com/Yamituki/go-review-cli/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/**
 * これらの変数はコマンドラインフラグの値を保持します。
 * rootCmd ルートコマンドを保持します。
 * cfgFile 設定ファイルのパスを保持します。
 * verbose 詳細出力フラグを保持します。
 */
var (
	rootCmd *cobra.Command
	cfgFile string
	verbose bool
)

// init 関数はパッケージの初期化を行います。
func init() {
	rootCmd = &cobra.Command{
		Use:     "go-review-cli",
		Short:   "Go Review CLI はコードレビューを支援するコマンドラインツールです。",
		Long:    `Go Review CLI は、Go言語で書かれたコードのレビューを効率化するための多機能なコマンドラインツールです。`,
		Version: version.GetVersion(),
	}
}

// Execute はルートコマンドを実行します。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

// initConfig は設定ファイルの初期化を行います。
func initConfig() {
	// viper の新しいインスタンスを作成します。
	v := viper.New()

	// Viperで設定ファイル読み込み（~/.go-review-cli/config.yaml）
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	}

	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.go-review-cli")

	if err := v.ReadInConfig(); err == nil {
		fmt.Println("設定ファイルを使用中:", v.ConfigFileUsed())
	}
}

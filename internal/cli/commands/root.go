package commands

import (
	"fmt"
	"os"

	"github.com/Yamituki/go-review-cli/internal/infrastructure/repository"
	"github.com/Yamituki/go-review-cli/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd *cobra.Command
	cfgFile string
	verbose bool
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd = &cobra.Command{
		Use:     "go-review-cli",
		Short:   "プロジェクトスキャフォールディングツール",
		Long:    `go-review-cli は、開発者がプロジェクトを素早く立ち上げるための対話型CLIツールです。統一されたプロジェクト構造、Git Flow管理、カスタムテンプレートサポートを提供します。`,
		Version: version.GetVersion(),
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "設定ファイルのパス（デフォルトは ~/.go-review-cli/config.yaml）")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "詳細な出力を有効にします")

	// 他のコマンドの初期化
	InitCreateCommand(rootCmd)
	InitConfigCommand(rootCmd)
}

// Execute はルートコマンドを実行します。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// initConfig は設定ファイルの初期化を行います。
func initConfig() {
	// viper の新しいインスタンスを作成します。
	v := viper.New()

	// 設定ファイルの有無を確認し、読み込みます。
	configRepo := repository.NewViperConfigRepository()
	if err := configRepo.EnsureConfigFile(); err != nil {
		fmt.Fprintln(os.Stderr, "設定ファイルの作成に失敗しました:", err)
		os.Exit(1)
	}

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

package commands

import (
	"github.com/Yamituki/go-review-cli/internal/infrastructure/di"
	"github.com/spf13/cobra"
)

// InitConfigCommand 設定の初期化コマンド
func InitConfigCommand(root *cobra.Command) {
	configCommand := &cobra.Command{
		Use:   "config",
		Short: "設定の管理する",
	}

	configGetCommand := &cobra.Command{
		Use:   "get",
		Short: "設定値を取得する",
		Args:  cobra.ExactArgs(1),
		Run:   runConfigGet,
	}

	configSetCommand := &cobra.Command{
		Use:   "set",
		Short: "設定値をセットする",
		Args:  cobra.ExactArgs(2),
		Run:   runConfigSet,
	}

	configListCommand := &cobra.Command{
		Use:   "list",
		Short: "設定値の一覧を表示する",
		Run:   runConfigList,
	}

	configResetCommand := &cobra.Command{
		Use:   "reset",
		Short: "設定値をリセットする",
		Run:   runConfigReset,
	}

	configCommand.AddCommand(configGetCommand)
	configCommand.AddCommand(configSetCommand)
	configCommand.AddCommand(configListCommand)
	configCommand.AddCommand(configResetCommand)

	root.AddCommand(configCommand)
}

// runConfigGet 設定値を取得する
func runConfigGet(cmd *cobra.Command, args []string) {
	// 依存性注入コンテナの初期化
	diContainer := di.NewContainer()

	// 設定リポジトリの取得
	configRepo := diContainer.GetConfigRepository()

	// キーの取得
	key := args[0]

	// 設定値の取得
	value, err := configRepo.Get(key)
	if err != nil {
		cmd.Println("設定が見つかりません:", err)
		return
	}

	// 設定値の表示
	cmd.Printf("%s: %s\n", key, value)
}

// runConfigSet 設定値をセットする
func runConfigSet(cmd *cobra.Command, args []string) {
	// 依存性注入コンテナの初期化
	diContainer := di.NewContainer()

	// 設定リポジトリの取得
	configRepo := diContainer.GetConfigRepository()

	// キーと値の取得
	key := args[0]
	value := args[1]

	// 設定値の設定
	err := configRepo.Set(key, value)
	if err != nil {
		cmd.Println("設定の更新に失敗しました:", err)
		return
	}

	// 成功メッセージの表示
	cmd.Printf("設定を更新しました: %s = %s\n", key, value)
}

// runConfigList 設定値の一覧を表示する
func runConfigList(cmd *cobra.Command, args []string) {
	// 依存性注入コンテナの初期化
	diContainer := di.NewContainer()

	// 設定リポジトリの取得
	configRepo := diContainer.GetConfigRepository()

	// 全ての設定値の取得
	settings, err := configRepo.GetAll()

	if err != nil {
		cmd.Println("設定の取得に失敗しました:", err)
		return
	}

	// 設定値の表示
	for key, value := range settings {
		cmd.Printf("%s: %s\n", key, value)
	}

}

// runConfigReset 設定値をリセットする
func runConfigReset(cmd *cobra.Command, args []string) {
	// 依存性注入コンテナの初期化
	diContainer := di.NewContainer()

	// 設定リポジトリの取得
	configRepo := diContainer.GetConfigRepository()

	// 設定値のリセット
	err := configRepo.Reset()
	if err != nil {
		cmd.Println("設定のリセットに失敗しました:", err)
		return
	}

	// 成功メッセージの表示
	cmd.Println("設定をリセットしました")
}

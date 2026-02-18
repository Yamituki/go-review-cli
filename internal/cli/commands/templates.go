package commands

import (
	"github.com/Yamituki/go-review-cli/internal/cli/prompts"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/di"
	"github.com/spf13/cobra"
)

// InitTemplateCommand テンプレート管理コマンドを初期化する
func InitTemplateCommand(root *cobra.Command) {
	// テンプレートコマンドを作成
	templatesCommand := &cobra.Command{
		Use:   "template",
		Short: "テンプレートを管理",
		Long:  "プロジェクトテンプレートの一覧表示、追加、削除、詳細表示を行います",
	}

	// テンプレートの一覧表示コマンドを作成
	templatesListCommand := &cobra.Command{
		Use:     "list",
		Short:   "テンプレートの一覧を表示",
		Aliases: []string{"ls"},
		Run:     runTemplateList,
	}

	// テンプレートの詳細表示コマンドを作成
	templatesShowCommand := &cobra.Command{
		Use:   "show [テンプレート名]",
		Short: "テンプレートの詳細を表示",
		Args:  cobra.ExactArgs(1),
		Run:   runTemplatesShow,
	}

	// テンプレートの追加コマンドを作成
	templatesAddCommand := &cobra.Command{
		Use:   "add [テンプレート名] [ソースパス]",
		Short: "テンプレートを追加",
		Args:  cobra.ExactArgs(2),
		Run:   runTemplateAdd,
	}

	// テンプレートの削除コマンドを作成
	templatesRemoveCommand := &cobra.Command{
		Use:     "remove [テンプレート名]",
		Short:   "テンプレートを削除",
		Aliases: []string{"rm", "delete"},
		Args:    cobra.ExactArgs(1),
		Run:     runTemplateRemove,
	}

	// サブコマンドを追加
	templatesCommand.AddCommand(templatesListCommand)
	templatesCommand.AddCommand(templatesShowCommand)
	templatesCommand.AddCommand(templatesAddCommand)
	templatesCommand.AddCommand(templatesRemoveCommand)

	// テンプレートコマンドをルートコマンドに追加
	root.AddCommand(templatesCommand)
}

// runTemplateList テンプレートの一覧を表示する関数
func runTemplateList(cmd *cobra.Command, args []string) {
	// コンテナ取得
	container := di.NewContainer()

	// リポジトリを取得
	templateRepo := container.GetTemplateRepository()

	// テンプレートの一覧を取得
	templates, err := templateRepo.List()
	if err != nil {
		cmd.Printf("テンプレートの取得に失敗しました: %v\n", err)
		return
	}

	// テンプレートの一覧を表示
	if len(templates) == 0 {
		cmd.Println("テンプレートが見つかりませんでした")
		return
	}

	cmd.Println("利用可能なテンプレート:")

	// テンプレートの一覧を表示
	for _, template := range templates {
		// 組み込みテンプレートかどうかを判断
		prefix := "[custom]"
		if templateRepo.IsBuiltin(template.Name) {
			prefix = "[built-in]"
		}

		cmd.Printf("  %s %-20s - %s\n", prefix, template.Name, template.Description)
	}
}

// runTemplatesShow テンプレートの詳細を表示する関数
func runTemplatesShow(cmd *cobra.Command, args []string) {
	// 引数チェック
	if len(args) == 0 {
		cmd.Println("テンプレート名を指定してください")
		return
	}

	// 名前引数の取得
	templateName := args[0]

	// コンテナ取得
	container := di.NewContainer()

	// リポジトリを取得
	templateRepo := container.GetTemplateRepository()

	// テンプレートの詳細を取得
	template, err := templateRepo.GetByName(templateName)
	if err != nil {
		cmd.Printf("テンプレートの取得に失敗しました: %v\n", err)
		return
	}

	// テンプレートの詳細を表示
	cmd.Printf("テンプレート名: %s\n", template.Name)
	cmd.Printf("バージョン: %s\n", template.Version)
	cmd.Printf("言語: %s\n", template.Language)
	cmd.Printf("タイプ: %s\n", template.Type)
	cmd.Printf("説明: %s\n", template.Description)
	cmd.Printf("パス: %s\n", template.Path)
}

// runTemplateAdd テンプレートを追加する関数
func runTemplateAdd(cmd *cobra.Command, args []string) {
	// 引数チェック
	if len(args) < 2 {
		cmd.Println("テンプレート名とソースパスを指定してください")
		return
	}

	// 名前とソースパスの取得
	templateName := args[0]
	sourcePath := args[1]

	// コンテナ取得
	container := di.NewContainer()

	// リポジトリを取得
	templateRepo := container.GetTemplateRepository()

	// テンプレートを追加
	if err := templateRepo.Add(templateName, sourcePath); err != nil {
		cmd.Printf("テンプレートの追加に失敗しました: %v\n", err)
		return
	}

	cmd.Printf("テンプレート '%s' を追加しました\n", templateName)
}

// runTemplateRemove テンプレートを削除する関数
func runTemplateRemove(cmd *cobra.Command, args []string) {
	// 引数チェック
	if len(args) == 0 {
		cmd.Println("テンプレート名を指定してください")
		return
	}

	// コンテナ取得
	container := di.NewContainer()

	// リポジトリを取得
	templateRepo := container.GetTemplateRepository()

	// テンプレート名の取得
	templateName := args[0]

	// テンプレート削除の確認
	confirm, err := prompts.PromptTemplateRemove(templateName)
	if err != nil {
		cmd.Printf("プロンプトの表示に失敗しました: %v\n", err)
		return
	}

	if !confirm {
		cmd.Println("テンプレートの削除をキャンセルしました")
		return
	}

	// テンプレートを削除
	if err := templateRepo.Remove(templateName); err != nil {
		cmd.Printf("テンプレートの削除に失敗しました: %v\n", err)
		return
	}

	cmd.Printf("テンプレート '%s' を削除しました\n", templateName)
}

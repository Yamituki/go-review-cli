package commands

import (
	"fmt"

	"github.com/Yamituki/go-review-cli/internal/application/usecase/dto"
	"github.com/Yamituki/go-review-cli/internal/cli/prompts"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/di"
	"github.com/spf13/cobra"
)

var (
	// グローバル変数定義
	createCmd *cobra.Command

	// パッケージ変数定義
	createName        string
	createType        string
	createModule      string
	createPath        string
	createDescription string
	createFramework   string
)

// InitCreateCommand createコマンドを初期化してrootコマンドに追加する
func InitCreateCommand(root *cobra.Command) {
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "新しいプロジェクトを作成",
		Long:  "指定されたパラメータで新しいGoプロジェクトを作成",
		RunE:  runCreate,
	}

	// 必須フラグの設定
	createCmd.Flags().StringVarP(&createName, "name", "n", "", "プロジェクト名")
	createCmd.Flags().StringVarP(&createType, "type", "t", "", "プロジェクトタイプ")
	createCmd.Flags().StringVarP(&createModule, "module", "m", "", "モジュール名")
	createCmd.Flags().StringVarP(&createPath, "path", "p", "", "プロジェクトパス")

	// オプションフラグの設定
	createCmd.Flags().StringVarP(&createDescription, "description", "d", "", "プロジェクトの説明")
	createCmd.Flags().StringVarP(&createFramework, "framework", "f", "", "使用するフレームワーク")

	// rootコマンドにcreateコマンドを追加
	root.AddCommand(createCmd)
}

// runCreate createコマンドの実行処理
func runCreate(cmd *cobra.Command, args []string) error {
	// 必須フラグが指定されていない場合、対話形式で入力を促す
	if createName == "" && createType == "" && createModule == "" && createPath == "" {
		// 対話形式でプロジェクト情報を取得
		prompt, err := prompts.PromptProjectCreation()
		if err != nil {
			return err
		}

		// 取得した情報を変数にセット
		createName = prompt.Name
		createType = prompt.Type
		createModule = prompt.Module
		createPath = prompt.Path
		createDescription = prompt.Description
		createFramework = prompt.Framework
	}

	// DIコンテナを作成
	container := di.NewContainer()

	// UseCaseの取得
	createUseCase := container.GetCreateProjectUseCase()

	// dtoの作成
	createProjectInput := dto.CreateProjectInput{
		Name:        createName,
		Type:        createType,
		Module:      createModule,
		Path:        createPath,
		Description: createDescription,
		Framework:   createFramework,
	}

	// UseCaseの実行
	createProjectOutput, err := createUseCase.Execute(createProjectInput)
	if err != nil {
		return err
	}

	// 成功メッセージの表示
	fmt.Println(createProjectOutput.Message)

	return nil
}

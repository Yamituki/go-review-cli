package usecase

import (
	"fmt"
	"testing"

	"github.com/Yamituki/go-review-cli/internal/application/usecase/dto"
	"github.com/Yamituki/go-review-cli/internal/domain/entity"
	"github.com/Yamituki/go-review-cli/internal/domain/service"
	"github.com/Yamituki/go-review-cli/internal/domain/value"
)

type MockProjectRepository struct {
	CreateFunc func(*entity.Project) error
}

// Create 新たなプロジェクトを作成するモック実装
func (m *MockProjectRepository) Create(project *entity.Project) error {
	return m.CreateFunc(project)
}

type MockTemplateRepository struct {
	GetByTypeFunc func(value.ProjectType) (*entity.Template, error)
}

// GetByType 指定されたタイプのプロジェクトを取得するモック実装
func (m *MockTemplateRepository) GetByType(projectType value.ProjectType) (*entity.Template, error) {
	return m.GetByTypeFunc(projectType)
}

// List すべてのプロジェクトをリストするモック実装
func (m *MockTemplateRepository) List() ([]*entity.Template, error) {
	// モックでは空のリストを返す
	return nil, nil
}

// Exists 指定されたパスにプロジェクトが存在するか確認するモック実装
func (m *MockProjectRepository) Exists(path string) (bool, error) {
	// モックでは常に存在しないと仮定
	return false, nil
}

type MockFileSystemService struct{}

// CreateDirectory 指定されたパスにディレクトリを作成するモック実装
func (m *MockFileSystemService) CreateDirectory(path string) error {
	return nil
}

// WriteFile 指定されたパスにファイルを書き込むモック実装
func (m *MockFileSystemService) WriteFile(path string, content string) error {
	// モックでは何もしない
	return nil
}

// ReadFile 指定されたパスからファイルを読み込むモック実装
func (m *MockFileSystemService) ReadFile(path string) (string, error) {
	// モックでは空の内容を返す
	return "", nil
}

// CopyDirectory 指定されたパスにディレクトリをコピーするモック実装
func (m *MockFileSystemService) CopyDirectory(src, dest string) error {
	// モックでは何もしない
	return nil
}

// DeleteFile 指定されたパスのファイルを削除するモック実装
func (m *MockFileSystemService) DeleteFile(path string) error {
	// モックでは何もしない
	return nil
}

// RenameFile 指定されたパスのファイルをリネームするモック実装
func (m *MockFileSystemService) RenameFile(oldPath, newPath string) error {
	// モックでは何もしない
	return nil
}

type MockGitService struct {
}

// Initialize 指定されたパスでGitリポジトリを初期化するモック実装
func (m *MockGitService) Initialize(path string) error {
	return nil
}

// CreateBranch 指定されたパスでGitブランチを作成するモック実装
func (m *MockGitService) CreateBranch(path, branchName string) error {
	return nil
}

// Commit 指定されたパスでGitコミットを作成するモック実装
func (m *MockGitService) Commit(path, message string) error {
	return nil
}

// SetupCommitMsgHook 指定されたパスでGitコミットメッセージフックを設定するモック実装
func (m *MockGitService) SetupCommitMsgHook(path string) error {
	return nil
}

// TestCreateProjectInteractor_Execute CreateProjectInteractorのExecuteメソッドをテスト
func TestCreateProjectInteractor_Execute(t *testing.T) {
	// input構造体
	type input struct {
		Name        string
		Type        string
		Framework   string
		Module      string
		Description string
		Path        string
	}

	// テストケース
	tests := []struct {
		name    string
		input   input
		wantErr bool
	}{
		{
			name: "正常にプロジェクトが作成される（APIタイプ、Frameworkあり）",
			input: input{
				Name:        "my-test-project",
				Type:        "api",
				Framework:   "gin",
				Module:      "github.com/user/my-test-project",
				Description: "自分のテストプロジェクトです",
				Path:        "/path/to/my-test-project",
			},
			wantErr: false,
		},

		{
			name: "正常にプロジェクトが作成される（CLIタイプ、Frameworkなし）",
			input: input{
				Name:        "my-cli-project",
				Type:        "cli",
				Framework:   "",
				Module:      "github.com/user/my-cli-project",
				Description: "自分のCLIプロジェクトです",
				Path:        "/path/to/my-cli-project",
			},
			wantErr: false,
		},

		{
			name: "異常系: entity.NewProjectでエラー（APIタイプでFrameworkなし）",
			input: input{
				Name:        "invalid-project",
				Type:        "api",
				Framework:   "",
				Module:      "github.com/user/invalid-project",
				Description: "無効なプロジェクトです",
				Path:        "/path/to/invalid-project",
			},
			wantErr: true,
		},

		{
			name: "異常系: projectRepo.Createでエラー（fmt.Errorfで\"repository error\"を返す）",
			input: input{
				Name:        "error-project",
				Type:        "cli",
				Framework:   "",
				Module:      "github.com/user/error-project",
				Description: "エラープロジェクトです",
				Path:        "/path/to/error-project",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックリポジトリの作成
			mockRepo := &MockProjectRepository{}

			// projectRepo.Createの動作を定義
			mockRepo.CreateFunc = func(project *entity.Project) error {
				if tt.input.Name == "error-project" {
					return fmt.Errorf("repository error")
				}
				return nil
			}

			// プロジェクトジェネレータの作成
			generator := service.NewProjectGenerator()

			// テンプレートプロセッサーの作成
			processor := service.NewTemplateProcessor()

			// テンプレートリポジトリの作成
			templateRepo := &MockTemplateRepository{}

			templateRepo.GetByTypeFunc = func(projectType value.ProjectType) (*entity.Template, error) {
				return entity.NewTemplate(
					"test-template",
					"1.0.0",
					"test",
					"go",
					"cli",
					"templates/test",
				)
			}

			// ファイルシステムサービスの作成
			fsService := &MockFileSystemService{}

			// Gitサービスの作成
			gitService := &MockGitService{}

			// インスタンス化
			interactor := NewCreateProjectInteractor(
				mockRepo,
				generator,
				processor,
				templateRepo,
				fsService,
				gitService,
			)

			// inputデータの準備
			input := dto.CreateProjectInput{
				Name:        tt.input.Name,
				Type:        tt.input.Type,
				Framework:   tt.input.Framework,
				Module:      tt.input.Module,
				Description: tt.input.Description,
				Path:        tt.input.Path,
			}

			// Executeメソッドの実行
			output, err := interactor.Execute(input)
			if (err != nil) != tt.wantErr {
				t.Errorf("テストケース %s において、期待されるエラー状態 %v と実際のエラー状態 %v が一致してない。エラー: %v", tt.name, tt.wantErr, err != nil, err)
			}

			// 異常系の処理
			if tt.wantErr {
				return
			}

			// 正常系の処理
			if !output.Success {
				t.Errorf("テストケース %s の処理で、出力が生成失敗した", tt.name)
			}

		})
	}

}

package usecase

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/Yamituki/go-review-cli/internal/application/usecase/dto"
	"github.com/Yamituki/go-review-cli/internal/domain/entity"
	"github.com/Yamituki/go-review-cli/internal/domain/repository"
	"github.com/Yamituki/go-review-cli/internal/domain/service"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/filesystem"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/git"
)

type CreateProjectInteractor struct {
	projectRepo       repository.ProjectRepository
	projectGenerator  *service.ProjectGenerator
	templateProcessor *service.TemplateProcessor
	templateRepo      repository.TemplateRepository
	fsService         filesystem.FileSystemService
	gitService        git.GitService
}

// NewCreateProjectInteractor CreateProjectInteractorインスタンスを生成します。
func NewCreateProjectInteractor(
	projectRepo repository.ProjectRepository,
	projectGenerator *service.ProjectGenerator,
	templateProcessor *service.TemplateProcessor,
	templateRepo repository.TemplateRepository,
	fsService filesystem.FileSystemService,
	gitService git.GitService,
) *CreateProjectInteractor {
	return &CreateProjectInteractor{
		projectRepo:       projectRepo,
		projectGenerator:  projectGenerator,
		templateProcessor: templateProcessor,
		templateRepo:      templateRepo,
		fsService:         fsService,
		gitService:        gitService,
	}
}

// Execute プロジェクト作成を実行
func (i *CreateProjectInteractor) Execute(input dto.CreateProjectInput) (*dto.CreateProjectOutput, error) {
	// entityの生成
	projectEntity, err := entity.NewProject(
		input.Name,
		input.Type,
		input.Framework,
		input.Module,
		input.Description,
		input.Path,
	)
	if err != nil {
		return nil, err
	}

	// プロジェクトの生成
	err = i.projectRepo.Create(projectEntity)
	if err != nil {
		return nil, err
	}

	// テンプレートを取得
	projectTemplate, err := i.templateRepo.GetByType(projectEntity.Type)
	if err != nil {
		return nil, err
	}

	// ディレクトリ構造を取得
	structure, err := i.projectGenerator.GenerateStructure(projectEntity)
	if err != nil {
		return nil, err
	}

	// ディレクトリ作成
	for _, dir := range structure {
		err := i.fsService.CreateDirectory(dir)
		if err != nil {
			return nil, err
		}
	}

	// テンプレート変数の準備
	variables := i.templateProcessor.PrepareVariables(projectEntity)

	// プロジェクトルートパスを作成
	projectRoot := filepath.Join(projectEntity.Path, projectEntity.Name.String())

	// テンプレートディレクトリからファイルをコピー
	err = i.fsService.CopyDirectory(projectTemplate.Path, projectRoot)
	if err != nil {
		return nil, err
	}

	// .tmplファイルの処理
	err = filepath.Walk(projectRoot, func(path string, info fs.FileInfo, err error) error {
		// エラーチェック
		if err != nil {
			return err
		}

		// ファイルのみを処理
		if info.IsDir() {
			return nil
		}

		// .tmplで終わるファイルのみを処理
		if !strings.HasSuffix(info.Name(), ".tmpl") {
			return nil
		}

		// ファイルを読み込む
		file, err := i.fsService.ReadFile(path)
		if err != nil {
			return err
		}

		// 変数置換
		processedContent, err := i.templateProcessor.Process(file, variables)
		if err != nil {
			return err
		}

		// 新しいファイル名を決定（.tmplを削除）
		newFileName := strings.TrimSuffix(info.Name(), ".tmpl")
		newFilePath := filepath.Join(filepath.Dir(path), newFileName)

		// 新しいファイルに書き込む
		err = i.fsService.WriteFile(newFilePath, processedContent)
		if err != nil {
			return err
		}

		// 元の.tmplファイルを削除
		err = i.fsService.DeleteFile(path)
		if err != nil {
			return err
		}

		return nil
	})

	// Gitリポジトリの初期化
	err = i.gitService.Initialize(projectRoot)
	if err != nil {
		return nil, err
	}

	// 初回コミット
	err = i.gitService.Commit(projectRoot, "🎉 chore: 初回コミット")
	if err != nil {
		return nil, err
	}

	// デベロップブランチの作成
	err = i.gitService.CreateBranch(projectRoot, "develop")
	if err != nil {
		return nil, err
	}

	// 成功処理
	out := &dto.CreateProjectOutput{
		Success:     true,
		Message:     "プロジェクトを作成しました",
		ProjectPath: fmt.Sprintf("%s/%s", projectEntity.Path, projectEntity.Name.String()),
	}

	return out, nil
}

package usecase

import (
	"fmt"

	"github.com/Yamituki/go-review-cli/internal/application/usecase/dto"
	"github.com/Yamituki/go-review-cli/internal/domain/entity"
	"github.com/Yamituki/go-review-cli/internal/domain/repository"
	"github.com/Yamituki/go-review-cli/internal/domain/service"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/filesystem"
)

type CreateProjectInteractor struct {
	projectRepo       repository.ProjectRepository
	projectGenerator  *service.ProjectGenerator
	templateProcessor *service.TemplateProcessor
	templateRepo      repository.TemplateRepository
	fsService         filesystem.FileSystemService
}

// NewCreateProjectInteractor CreateProjectInteractorインスタンスを生成します。
func NewCreateProjectInteractor(
	projectRepo repository.ProjectRepository,
	projectGenerator *service.ProjectGenerator,
	templateProcessor *service.TemplateProcessor,
	templateRepo repository.TemplateRepository,
	fsService filesystem.FileSystemService,
) *CreateProjectInteractor {
	return &CreateProjectInteractor{
		projectRepo:       projectRepo,
		projectGenerator:  projectGenerator,
		templateProcessor: templateProcessor,
		templateRepo:      templateRepo,
		fsService:         fsService,
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

	// 成功処理
	out := &dto.CreateProjectOutput{
		Success:     true,
		Message:     "プロジェクトを作成しました",
		ProjectPath: fmt.Sprintf("%s/%s", projectEntity.Path, projectEntity.Name.String()),
	}

	return out, nil
}

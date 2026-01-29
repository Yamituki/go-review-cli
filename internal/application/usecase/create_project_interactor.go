package usecase

import (
	"fmt"

	"github.com/Yamituki/go-review-cli/internal/application/usecase/dto"
	"github.com/Yamituki/go-review-cli/internal/domain/entity"
	"github.com/Yamituki/go-review-cli/internal/domain/repository"
	"github.com/Yamituki/go-review-cli/internal/domain/service"
)

type CreateProjectInteractor struct {
	projectRepo       repository.ProjectRepository
	projectGenerator  *service.ProjectGenerator
	templateProcessor *service.TemplateProcessor
}

// NewCreateProjectInteractor CreateProjectInteractorインスタンスを生成します。
func NewCreateProjectInteractor(
	projectRepo repository.ProjectRepository,
	projectGenerator *service.ProjectGenerator,
	templateProcessor *service.TemplateProcessor,
) *CreateProjectInteractor {
	return &CreateProjectInteractor{
		projectRepo:       projectRepo,
		projectGenerator:  projectGenerator,
		templateProcessor: templateProcessor,
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

	// 成功処理
	out := &dto.CreateProjectOutput{
		Success:     true,
		Message:     "プロジェクトを作成しました",
		ProjectPath: fmt.Sprintf("%s/%s", projectEntity.Path, projectEntity.Name.String()),
	}

	return out, nil
}

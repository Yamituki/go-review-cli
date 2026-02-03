package di

import (
	"github.com/Yamituki/go-review-cli/internal/application/usecase"
	"github.com/Yamituki/go-review-cli/internal/domain/service"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/repository"
)

type Container struct {
	createProjectUseCase usecase.CreateProjectUseCase
}

// NewContainer コンテナの新規作成
func NewContainer() *Container {
	// UseCaseの初期化
	projectRepo := repository.NewFileSystemProjectRepository()
	projectGenerator := service.NewProjectGenerator()
	templateProcessor := service.NewTemplateProcessor()

	return &Container{
		createProjectUseCase: usecase.NewCreateProjectInteractor(projectRepo, projectGenerator, templateProcessor),
	}
}

// GetCreateProjectUseCase CreateProjectUseCaseの取得
func (c *Container) GetCreateProjectUseCase() usecase.CreateProjectUseCase {
	return c.createProjectUseCase
}

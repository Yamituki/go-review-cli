package di

import (
	"github.com/Yamituki/go-review-cli/internal/application/usecase"
	domainRepository "github.com/Yamituki/go-review-cli/internal/domain/repository"
	"github.com/Yamituki/go-review-cli/internal/domain/service"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/filesystem"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/git"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/repository"
)

type Container struct {
	createProjectUseCase usecase.CreateProjectUseCase
	configRepository     domainRepository.ConfigRepository
}

// NewContainer コンテナの新規作成
func NewContainer() *Container {
	// プロジェクトリポジトリの作成
	projectRepo := repository.NewFileSystemProjectRepository()
	// プロジェクトジェネレータの作成
	projectGenerator := service.NewProjectGenerator()
	// テンプレートプロセッサの作成
	templateProcessor := service.NewTemplateProcessor()
	// テンプレートリポジトリの作成
	templateRepo := repository.NewFileSystemTemplateRepository()
	// ファイルシステムサービスの作成
	fsService := filesystem.NewAferoFileSystemService()
	// GoGitサービスの作成
	gitService := git.NewGoGitService()

	// configリポジトリの作成
	configRepository := repository.NewViperConfigRepository()

	return &Container{
		createProjectUseCase: usecase.NewCreateProjectInteractor(
			projectRepo,
			projectGenerator,
			templateProcessor,
			templateRepo,
			fsService,
			gitService,
		),
		configRepository: configRepository,
	}
}

// GetCreateProjectUseCase CreateProjectUseCaseの取得
func (c *Container) GetCreateProjectUseCase() usecase.CreateProjectUseCase {
	return c.createProjectUseCase
}

// GetConfigRepository ConfigRepositoryの取得
func (c *Container) GetConfigRepository() domainRepository.ConfigRepository {
	return c.configRepository
}

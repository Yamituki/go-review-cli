package usecase

import "github.com/Yamituki/go-review-cli/internal/application/usecase/dto"

type CreateProjectUseCase interface {
	// Execute プロジェクト作成を実行
	Execute(input dto.CreateProjectInput) (*dto.CreateProjectOutput, error)
}

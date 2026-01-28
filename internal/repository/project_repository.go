package repository

import "github.com/Yamituki/go-review-cli/internal/domain/entity"

type ProjectRepository interface {
	// Create 新たなプロジェクトを作成する
	Create(project *entity.Project) error
	// Exists 指定されたパスにプロジェクトが存在するか確認する
	Exists(path string) (bool, error)
}
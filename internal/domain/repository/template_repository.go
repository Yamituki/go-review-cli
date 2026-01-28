package repository

import (
	"github.com/Yamituki/go-review-cli/internal/domain/entity"
	"github.com/Yamituki/go-review-cli/internal/domain/value"
)

type TemplateRepository interface {
	// GetByType 指定されたタイプのテンプレートを取得する
	GetByType(projectType value.ProjectType) (*entity.Template, error)
	// List 利用可能なテンプレートの一覧を取得する
	List() ([]*entity.Template, error)
}

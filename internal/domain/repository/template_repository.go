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
	// GetByName 指定された名前のテンプレートを取得する
	GetByName(name string) (*entity.Template, error)
	// Add 新しいテンプレートを追加する
	Add(name, sourcePath string) error
	// Remove 指定されたテンプレートを削除する
	Remove(name string) error
	// IsBuiltin 指定されたテンプレートが組み込みテンプレートかどうかを判断する
	IsBuiltin(name string) bool
}

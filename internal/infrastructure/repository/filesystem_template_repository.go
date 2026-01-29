package repository

import (
	"fmt"

	"github.com/Yamituki/go-review-cli/internal/domain/entity"
	"github.com/Yamituki/go-review-cli/internal/domain/value"
)

type FileSystemTemplateRepository struct{}

// NewFileSystemTemplateRepository 新たなFileSystemTemplateRepositoryを生成します。
func NewFileSystemTemplateRepository() *FileSystemTemplateRepository {
	return &FileSystemTemplateRepository{}
}

// GetByType テンプレートタイプに基づいてテンプレートを取得します。
func (r *FileSystemTemplateRepository) GetByType(projectType value.ProjectType) (*entity.Template, error) {
	return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: GetByTypeは未実装です")
}

// List すべてのテンプレートをリストします。
func (r *FileSystemTemplateRepository) List() ([]*entity.Template, error) {
	return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: Listは未実装です")
}

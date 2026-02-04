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
	// テンプレートの初期化
	var template *entity.Template
	var err error

	// タイプに基づいてテンプレートを選択
	switch projectType {
	case value.ProjectTypeAPI:
		template, err = entity.NewTemplate("go-api", "1.0.0", "Go RESTful API template", "go", "api", "templates/go-api")
		if err != nil {
			return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの生成に失敗しました: %w", err)
		}
	case value.ProjectTypeCLI:
		template, err = entity.NewTemplate("go-cli", "1.0.0", "Go CLI tool template", "go", "cli", "templates/go-cli")
		if err != nil {
			return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの生成に失敗しました: %w", err)
		}
	case value.ProjectTypeMicroservice:
		template, err = entity.NewTemplate("go-microservice", "1.0.0", "Go microservice template", "go", "microservice", "templates/go-microservice")
		if err != nil {
			return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの生成に失敗しました: %w", err)
		}
	default:
		return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: 未対応のプロジェクトタイプ %s", projectType.String())
	}

	return template, nil
}

// List すべてのテンプレートをリストします。
func (r *FileSystemTemplateRepository) List() ([]*entity.Template, error) {
	// テンプレートのリストを初期化
	templates := []*entity.Template{}

	// すべてのプロジェクトタイプを列挙
	projectTypes := []value.ProjectType{
		value.ProjectTypeAPI,
		value.ProjectTypeCLI,
		value.ProjectTypeMicroservice,
	}

	// 各プロジェクトタイプに対してテンプレートを取得
	for _, projectType := range projectTypes {
		template, err := r.GetByType(projectType)
		if err != nil {
			return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの取得に失敗しました: %w", err)
		}
		templates = append(templates, template)
	}

	return templates, nil
}

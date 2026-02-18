package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yamituki/go-review-cli/internal/domain/entity"
	"github.com/Yamituki/go-review-cli/internal/domain/value"
	"github.com/Yamituki/go-review-cli/internal/infrastructure/filesystem"
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

	// カスタムテンプレートの一覧を取得
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: ホームディレクトリの取得に失敗しました: %w", err)
	}

	customTemplatesDir := filepath.Join(homeDir, ".go-review-cli", "templates")

	entries, err := os.ReadDir(customTemplatesDir)
	if err != nil {
		if os.IsNotExist(err) {
			return templates, nil // ディレクトリ未作成の場合はスキップ
		}

		return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: カスタムテンプレートディレクトリの読み取りに失敗しました: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			templateName := entry.Name()
			templatePath := filepath.Join(customTemplatesDir, templateName)
			template, err := entity.NewTemplate(templateName, "1.0.0", "User-defined template", "go", templateName, templatePath)
			if err != nil {
				return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: カスタムテンプレートの生成に失敗しました: %w", err)
			}
			templates = append(templates, template)
		}
	}

	return templates, nil
}

// GetByName テンプレート名に基づいてテンプレートを取得します。
func (r *FileSystemTemplateRepository) GetByName(name string) (*entity.Template, error) {
	// テンプレート名プレフィックス処理
	templateName := r.normalizeTemplateName(name)

	// 組み込みテンプレートかどうかを判断
	if isbuiltin := r.IsBuiltin(name); isbuiltin {
		return r.GetByType(value.ProjectType(templateName))
	}

	// カスタムテンプレートのパスを取得
	templatePath, err := r.getCustomTemplatePath(templateName)
	if err != nil {
		return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: カスタムテンプレートのパスの取得に失敗しました: %w", err)
	}

	// 存在チェック
	exists, err := r.customTemplateExists(templateName)
	if err != nil {
		return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの存在チェックに失敗しました: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートが見つかりません: %s", name)
	}

	// テンプレートの詳細を取得
	template, err := entity.NewTemplate(name, "1.0.0", "User-defined template", "go", templateName, templatePath)
	if err != nil {
		return nil, fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの生成に失敗しました: %w", err)
	}

	return template, nil
}

// Add 新しいテンプレートを追加
func (r *FileSystemTemplateRepository) Add(name, sourcePath string) error {
	// テンプレート名プレフィックス処理
	templateName := r.normalizeTemplateName(name)

	// 組み込みテンプレートかどうかを判断
	if isbuiltin := r.IsBuiltin(name); isbuiltin {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: 組み込みテンプレートは追加できません: %s", name)
	}

	// カスタムテンプレートと同名チェック
	exists, err := r.customTemplateExists(templateName)
	if err != nil {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの存在チェックに失敗しました: %w", err)
	}

	if exists {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: 同名のテンプレートが既に存在します: %s", name)
	}

	// テンプレートソースパスの存在チェック
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートパスが存在しません: %s", sourcePath)
	}

	// カスタムテンプレートのパスを取得
	templatePath, err := r.getCustomTemplatePath(templateName)
	if err != nil {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: カスタムテンプレートのパスの取得に失敗しました: %w", err)
	}

	// テンプレートパスのディレクトリを作成
	if err := os.MkdirAll(filepath.Dir(templatePath), os.ModePerm); err != nil {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートディレクトリの作成に失敗しました: %w", err)
	}

	// ファイルシステムを取得
	fsService := filesystem.NewAferoFileSystemService()

	// テンプレートをコピー
	if err := fsService.CopyDirectory(sourcePath, templatePath); err != nil {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートのコピーに失敗しました: %w", err)
	}

	return nil
}

// Remove テンプレートを削除します（組み込みテンプレートは削除できません）。
func (r *FileSystemTemplateRepository) Remove(name string) error {
	// テンプレート名プレフィックス処理
	templateName := r.normalizeTemplateName(name)

	// 組み込みテンプレートかどうかを判断
	if isbuiltin := r.IsBuiltin(name); isbuiltin {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: 組み込みテンプレートは削除できません: %s", name)
	}

	// カスタムテンプレートのパスを取得
	templatePath, err := r.getCustomTemplatePath(templateName)
	if err != nil {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: カスタムテンプレートのパスの取得に失敗しました: %w", err)
	}

	// 存在チェック
	exists, err := r.customTemplateExists(templateName)
	if err != nil {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの存在チェックに失敗しました: %w", err)
	}
	if !exists {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートが見つかりません: %s", name)
	}

	// テンプレートを削除
	if err := os.RemoveAll(templatePath); err != nil {
		return fmt.Errorf("ファイルシステムテンプレートリポジトリ: テンプレートの削除に失敗しました: %w", err)
	}

	return nil
}

// IsBuiltin 組み込みテンプレートかどうかを判断します。
func (r *FileSystemTemplateRepository) IsBuiltin(name string) bool {
	// テンプレート名プレフィックス処理
	templateName := r.normalizeTemplateName(name)

	// 組み込みテンプレートかどうかを判断
	_, err := value.NewProjectType(templateName)
	if err != nil {
		return false
	}

	return true
}

// normalizeTemplateName テンプレート名を正規化します。
func (r *FileSystemTemplateRepository) normalizeTemplateName(name string) string {
	return strings.TrimPrefix(name, "go-")
}

// getCustomTemplatePath カスタムテンプレートのパスを取得します。
func (r *FileSystemTemplateRepository) getCustomTemplatePath(name string) (string, error) {
	// ホームディレクトリ取得
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("ファイルシステムテンプレートリポジトリ: ホームディレクトリの取得に失敗しました: %w", err)
	}

	// テンプレートのパスを構築
	templatePath := filepath.Join(homeDir, ".go-review-cli", "templates", name)

	return templatePath, nil
}

// customTemplateExists カスタムテンプレートが存在するかどうかをチェックします。
func (r *FileSystemTemplateRepository) customTemplateExists(name string) (bool, error) {
	templatePath, err := r.getCustomTemplatePath(name)
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

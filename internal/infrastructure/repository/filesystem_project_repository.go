package repository

import (
	"errors"
	"fmt"
	"os"

	"github.com/Yamituki/go-review-cli/internal/domain/entity"
)

type FileSystemProjectRepository struct{}

// NewFileSystemProjectRepository 新たなFileSystemProjectRepositoryを生成します。
func NewFileSystemProjectRepository() *FileSystemProjectRepository {
	return &FileSystemProjectRepository{}
}

// Create プロジェクトをファイルシステムに作成します。
func (r *FileSystemProjectRepository) Create(project *entity.Project) error {
	// rootディレクトリのパスを生成
	rootPath := fmt.Sprintf("%s/%s", project.Path, project.Name.String())

	// 存在チェック
	exists, err := r.Exists(rootPath)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("プロジェクトが既に存在します: %s", rootPath)
	}

	// ディレクトリ作成
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		return fmt.Errorf("プロジェクトのディレクトリ作成に失敗しました: %w", err)
	}

	return nil
}

// Exists プロジェクトがファイルシステムに存在するか確認します。
func (r *FileSystemProjectRepository) Exists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}

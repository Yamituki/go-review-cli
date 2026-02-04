package filesystem

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/spf13/afero"
)

type AferoFileSystemService struct {
	fs afero.Fs
}

// NewAferoFileSystemService AferoFileSystemServiceのコンストラクタ
func NewAferoFileSystemService() *AferoFileSystemService {
	// AferoのOsFsを使用してファイルシステムを初期化
	fs := afero.NewOsFs()

	return &AferoFileSystemService{
		fs: fs,
	}
}

// CreateDirectory ディレクトリを作成します。
func (afs *AferoFileSystemService) CreateDirectory(path string) error {
	// ディレクトリを作成
	if err := afs.fs.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("Aferoファイルシステムサービス: ディレクトリの作成に失敗しました: %w", err)
	}

	return nil
}

// WriteFile ファイルにデータを書き込みます。
func (afs *AferoFileSystemService) WriteFile(path string, content string) error {
	// ファイルにデータを書き込み
	if err := afero.WriteFile(afs.fs, path, []byte(content), 0644); err != nil {
		return fmt.Errorf("Aferoファイルシステムサービス: ファイルの書き込みに失敗しました: %w", err)
	}

	return nil
}

// ReadFile ファイルからデータを読み込みます。
func (afs *AferoFileSystemService) ReadFile(path string) (string, error) {
	// ファイルからデータを読み込み
	data, err := afero.ReadFile(afs.fs, path)
	if err != nil {
		return "", fmt.Errorf("Aferoファイルシステムサービス: ファイルの読み込みに失敗しました: %w", err)
	}

	return string(data), nil
}

// CopyDirectory ディレクトリをコピーします。
func (afs *AferoFileSystemService) CopyDirectory(src, dest string) error {
	// ディレクトリのコピー元を確認
	_, err := afs.fs.Stat(src)
	if err != nil {
		return fmt.Errorf("Aferoファイルシステムサービス: コピー元ディレクトリの情報取得に失敗しました: %w", err)
	}

	// コピー先ディレクトリを作成
	if err := afs.CreateDirectory(dest); err != nil {
		return fmt.Errorf("Aferoファイルシステムサービス: コピー先ディレクトリの作成に失敗しました: %w", err)
	}

	// コピー元走査
	err = afero.Walk(afs.fs, src, func(path string, info fs.FileInfo, err error) error {
		// エラーチェック
		if err != nil {
			return err
		}

		// 相対パスを計算（srcからの相対パス）
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// コピー先のパスを生成
		destPath := filepath.Join(dest, relPath)

		// ディレクトリの場合
		if info.IsDir() {
			// コピー先にディレクトリを作成
			return afs.CreateDirectory(destPath)
		}

		// ファイルの場合
		content, err := afs.ReadFile(path)
		if err != nil {
			return err
		}

		// コピー先にファイルを書き込み
		return afs.WriteFile(destPath, content)

	})

	if err != nil {
		return fmt.Errorf("Aferoファイルシステムサービス: ディレクトリのコピーに失敗しました: %w", err)
	}

	return nil
}

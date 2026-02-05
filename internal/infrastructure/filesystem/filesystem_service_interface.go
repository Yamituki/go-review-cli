package filesystem

type FileSystemService interface {
	// CreateDirectory ディレクトリを作成
	CreateDirectory(path string) error
	// WriteFile ファイルにデータを書き込む
	WriteFile(path string, content string) error
	// ReadFile ファイルからデータを読み込む
	ReadFile(path string) (string, error)
	// CopyDirectory ディレクトリをコピー
	CopyDirectory(src, dest string) error
	// DeleteFile ファイルを削除
	DeleteFile(path string) error
	// RenameFile ファイルをリネーム
	RenameFile(oldPath, newPath string) error
}

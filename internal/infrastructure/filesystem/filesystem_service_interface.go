package filesystem

type FileSystemService interface {
	// CreateDirectory ディレクトリを作成する
	CreateDirectory(path string) error
	// WriteFile ファイルにデータを書き込む
	WriteFile(path string, content string) error
	// ReadFile ファイルからデータを読み込む
	ReadFile(path string) (string, error)
	// CopyDirectory ディレクトリをコピーする
	CopyDirectory(src, dest string) error
}

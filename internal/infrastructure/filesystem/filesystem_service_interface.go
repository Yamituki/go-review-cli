package filesystem

type FileSystemService interface {
	// CreateDirectory ディレクトリを作成します。
	CreateDirectory(path string) error
	// WriteFile ファイルにデータを書き込みます。
	WriteFile(path string, content string) error
	// ReadFile ファイルからデータを読み込みます。
	ReadFile(path string) (string, error)
	// CopyDirectory ディレクトリをコピーします。
	CopyDirectory(src, dest string) error
}

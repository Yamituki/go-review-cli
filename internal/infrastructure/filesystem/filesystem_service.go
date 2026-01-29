package filesystem

import "fmt"

type AferoFileSystemService struct{}

// NewAferoFileSystemService AferoFileSystemServiceのコンストラクタ
func NewAferoFileSystemService() *AferoFileSystemService {
	return &AferoFileSystemService{}
}

// CreateDirectory ディレクトリを作成します。
func (afs *AferoFileSystemService) CreateDirectory(path string) error {
	return fmt.Errorf("Aferoファイルシステムサービス: CreateDirectoryは未実装です")
}

// WriteFile ファイルにデータを書き込みます。
func (afs *AferoFileSystemService) WriteFile(path string, content string) error {
	return fmt.Errorf("Aferoファイルシステムサービス: WriteFileは未実装です")
}

// ReadFile ファイルからデータを読み込みます。
func (afs *AferoFileSystemService) ReadFile(path string) (string, error) {
	return "", fmt.Errorf("Aferoファイルシステムサービス: ReadFileは未実装です")
}

// CopyDirectory ディレクトリをコピーします。
func (afs *AferoFileSystemService) CopyDirectory(src, dest string) error {
	return fmt.Errorf("Aferoファイルシステムサービス: CopyDirectoryは未実装です")
}

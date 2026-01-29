package git

import "fmt"

type GoGitService struct{}

func NewGoGitService() *GoGitService {
	return &GoGitService{}
}

// Initialize Gitリポジトリを初期化します。
func (g *GoGitService) Initialize(path string) error {
	return fmt.Errorf("Gitサービス: Initializeは未実装です")
}

// CreateBranch Gitブランチを作成します。
func (g *GoGitService) CreateBranch(path, branchName string) error {
	return fmt.Errorf("Gitサービス: CreateBranchは未実装です")
}

// Commit Gitコミットを作成します。
func (g *GoGitService) Commit(path, message string) error {
	return fmt.Errorf("Gitサービス: Commitは未実装です")
}

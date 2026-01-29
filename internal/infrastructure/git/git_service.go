package git

import "fmt"

type GoGitService struct{}

func NewGoGitService() *GoGitService {
	return &GoGitService{}
}

// 初期化
func (g *GoGitService) Initialize(path string) error {
	return fmt.Errorf("Gitサービス: Initializeは未実装です")
}

// ブランチの作成
func (g *GoGitService) CreateBranch(path, branchName string) error {
	return fmt.Errorf("Gitサービス: CreateBranchは未実装です")
}

// コミット
func (g *GoGitService) Commit(path, message string) error {
	return fmt.Errorf("Gitサービス: Commitは未実装です")
}

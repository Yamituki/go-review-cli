package git

type GitService interface {
	// 初期化
	Initialize(path string) error
	// ブランチの作成
	CreateBranch(path, branchName string) error
	// コミット
	Commit(path, message string) error
}

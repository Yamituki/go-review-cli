package git

type GitService interface {
	// Initialize Gitリポジトリを初期化します。
	Initialize(path string) error
	// CreateBranch Gitブランチを作成します。
	CreateBranch(path, branchName string) error
	// Commit Gitコミットを作成します。
	Commit(path, message string) error
	// SetupCommitMsgHook Gitコミットメッセージフックを設定します。
	SetupCommitMsgHook(path string) error
}

package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type GoGitService struct{}

func NewGoGitService() *GoGitService {
	return &GoGitService{}
}

// Initialize Gitリポジトリを初期化します。
func (g *GoGitService) Initialize(path string) error {
	// リポジトリを初期化
	_, err := git.PlainInit(path, false)
	if err != nil {
		return fmt.Errorf("Gitサービス: Initializeエラー: %w", err)
	}

	return nil
}

// CreateBranch Gitブランチを作成します。
func (g *GoGitService) CreateBranch(path, branchName string) error {
	// リポジトリを開く
	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("Gitサービス: リポジトリオープンエラー: %w", err)
	}

	// HEADの参照を取得
	headRef, err := repo.Head()
	if err != nil {
		return fmt.Errorf("Gitサービス: HEAD取得エラー: %w", err)
	}

	// ブランチ参照名を作成
	branchNameRef := plumbing.NewBranchReferenceName(branchName)

	// ハッシュ参照を作成
	ref := plumbing.NewHashReference(branchNameRef, headRef.Hash())

	// 参照を設定
	if err := repo.Storer.SetReference(ref); err != nil {
		return fmt.Errorf("Gitサービス: ブランチ作成エラー: %w", err)
	}

	return nil
}

// Commit Gitコミットを作成します。
func (g *GoGitService) Commit(path, message string) error {
	// リポジトリを開く
	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("Gitサービス: リポジトリオープンエラー: %w", err)
	}

	// ワークツリーを取得
	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("Gitサービス: ワークツリー取得エラー: %w", err)
	}

	// すべての変更をステージング
	if err := worktree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		return fmt.Errorf("Gitサービス: 変更ステージングエラー: %w", err)
	}

	// コミットを作成
	_, err = worktree.Commit(message, &git.CommitOptions{})
	if err != nil {
		return fmt.Errorf("Gitサービス: コミットエラー: %w", err)
	}

	return nil
}

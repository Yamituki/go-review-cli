package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Yamituki/go-review-cli/internal/templates"
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

// SetupCommitMsgHook Gitコミットメッセージフックを設定します。
func (g *GoGitService) SetupCommitMsgHook(path string) error {
	// パスを確認
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("Gitサービス: 指定パスが存在しません: %s", path)
	}

	// フックディレクトリのパスを作成
	hookDir := filepath.Join(path, ".git", "hooks")

	// コピー先のディレクトリを作成
	if _, err := os.Stat(hookDir); os.IsNotExist(err) {
		if err := os.MkdirAll(hookDir, 0755); err != nil {
			return fmt.Errorf("Gitサービス: フックディレクトリ作成エラー: %w", err)
		}
	}

	// prepare-commit-msgフックのパスを作成
	hookPath := filepath.Join(hookDir, "prepare-commit-msg")

	// フックファイルを書き込み
	if err := os.WriteFile(hookPath, []byte(templates.PrepareCommitMsgHook), 0755); err != nil {
		return fmt.Errorf("Gitサービス: フックファイル書き込みエラー: %w", err)
	}

	// 実行権限chmod+xを設定
	if err := os.Chmod(hookPath, 0755); err != nil {
		return fmt.Errorf("Gitサービス: フックファイル権限設定エラー: %w", err)
	}

	return nil
}

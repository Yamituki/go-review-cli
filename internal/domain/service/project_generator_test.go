package service

import (
	"testing"

	"github.com/Yamituki/go-review-cli/internal/domain/entity"
)

// TestGenerateStructure ディレクトリ構造生成のテスト
func TestGenerateStructure(t *testing.T) {
	// テストケース
	testCases := []struct {
		name     string
		project  *entity.Project
		wantDirs []string
		wantErr  bool
	}{}

	// 正常系：基本的なディレクトリ構造
	baseProject, err := entity.NewProject(
		"test-project",
		"cli",
		"",
		"github.com/test/test-project",
		"これはテストプロジェクトです",
		"/tmp",
	)

	if err != nil {
		t.Fatal("正常系：基本的なディレクトリ構造のプロジェクト生成に失敗しました:", err)
	}

	testCases = append(testCases, struct {
		name     string
		project  *entity.Project
		wantDirs []string
		wantErr  bool
	}{
		name:    "正常系：基本的なディレクトリ構造",
		project: baseProject,
		wantDirs: []string{
			"root",
			"cmd",
			"internal",
			"internal/domain",
			"internal/usecase",
			"internal/handler",
			"internal/infrastructure",
			"pkg",
			"test",
			"configs",
			"scripts",
			"docs",
		},
		wantErr: false,
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			generator := NewProjectGenerator()
			gotStructure, err := generator.GenerateStructure(tc.project)

			if (err != nil) != tc.wantErr {
				t.Fatalf("%s: エラーの有無が期待値と異なります。 got: %v", tc.name, err)
			}

			// 期待されるディレクトリがすべて存在するか確認
			for _, dirKey := range tc.wantDirs {
				if _, exists := gotStructure[dirKey]; !exists {
					t.Errorf("%s: 期待されるディレクトリ '%s' が生成されていません", tc.name, dirKey)
				}
			}
		})
	}
}

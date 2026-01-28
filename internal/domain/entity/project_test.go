package entity

import (
	"testing"
)

// TestNewProject は NewProject 関数のテストを行います。
func TestNewProject(t *testing.T) {
	// input構造体
	type input struct {
		Name        string
		Type        string
		Framework   string
		Module      string
		Description string
		Path        string
	}

	// テストケース
	tests := []struct {
		name    string
		input   input
		wantErr bool
	}{
		{
			name: "APIタイプでフレームワークが指定されている場合、正常にプロジェクトが生成される",
			input: input{
				Name:        "my-test-project",
				Type:        "api",
				Framework:   "gin",
				Module:      "github.com/user/my-test-project",
				Description: "自分のテストプロジェクトです",
				Path:        "/path/to/my-test-project",
			},
			wantErr: false,
		},

		{
			name: "CLIタイプでフレームワークが空文字の場合、正常にプロジェクトが生成される",
			input: input{
				Name:        "my-cli-project",
				Type:        "cli",
				Framework:   "",
				Module:      "github.com/user/my-cli-project",
				Description: "自分のCLIプロジェクトです",
				Path:        "/path/to/my-cli-project",
			},
			wantErr: false,
		},

		{
			name: "APIタイプでフレームワークが空文字の場合、エラーを返す",
			input: input{
				Name:        "invalid-project",
				Type:        "api",
				Framework:   "",
				Module:      "github.com/user/invalid-project",
				Description: "無効なプロジェクトです",
				Path:        "/path/to/invalid-project",
			},
			wantErr: true,
		},

		{
			name: "CLIタイプでフレームワークが指定されている場合、エラーを返す",
			input: input{
				Name:        "invalid-cli-project",
				Type:        "cli",
				Framework:   "echo",
				Module:      "github.com/user/invalid-cli-project",
				Description: "無効なCLIプロジェクトです",
				Path:        "/path/to/invalid-cli-project",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProject(
				tt.input.Name,
				tt.input.Type,
				tt.input.Framework,
				tt.input.Module,
				tt.input.Description,
				tt.input.Path,
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("プロジェクトの生成に失敗しました。 got error: %v, want error: %v", err, tt.wantErr)
				return
			}

		})
	}

}

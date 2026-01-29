package service

import (
	"testing"

	"github.com/Yamituki/go-review-cli/internal/domain/entity"
)

// TestPrepareVariables テンプレート変数の準備をテストします。
func TestPrepareVariables(t *testing.T) {
	// テストケース
	testCases := []struct {
		name          string
		project       *entity.Project
		wantVariables map[string]string
	}{}

	// 正常系：フレームワークあり

	// APIプロジェクト
	apiProject, error := entity.NewProject(
		"test-project",
		"api",
		"gin",
		"github.com/user/test-project",
		"Test project",
		"/tmp",
	)

	if error != nil {
		t.Fatalf("正常系：フレームワークありのAPIプロジェクト生成に失敗しました: %v", error)
	}

	testCases = append(testCases, struct {
		name          string
		project       *entity.Project
		wantVariables map[string]string
	}{
		name:    "正常系：フレームワークありのAPIプロジェクト",
		project: apiProject,
		wantVariables: map[string]string{
			"ProjectName": "test-project",
			"ModuleName":  "github.com/user/test-project",
			"ProjectType": "api",
			"Framework":   "gin",
			"Description": "Test project",
		},
	})

	// 正常系：フレームワークなし

	// CLIプロジェクト
	cliProject, error := entity.NewProject(
		"test-cli",
		"cli",
		"",
		"github.com/user/test-cli",
		"Test CLI",
		"/tmp",
	)

	if error != nil {
		t.Fatalf("正常系：フレームワークなしのCLIプロジェクト生成に失敗しました: %v", error)
	}

	testCases = append(testCases, struct {
		name          string
		project       *entity.Project
		wantVariables map[string]string
	}{
		name:    "正常系：フレームワークなしのCLIプロジェクト",
		project: cliProject,
		wantVariables: map[string]string{
			"ProjectName": "test-cli",
			"ModuleName":  "github.com/user/test-cli",
			"ProjectType": "cli",
			"Framework":   "",
			"Description": "Test CLI",
		},
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processor := NewTemplateProcessor()
			gotVariables := processor.PrepareVariables(tc.project)
			for key, wantValue := range tc.wantVariables {
				gotValue, exists := gotVariables[key]

				if !exists {
					t.Fatalf("変数 %s が存在しません", key)
				}

				if gotValue != wantValue {
					t.Errorf("変数 %s の値が期待値と異なります。期待値: %s, 実際の値: %s", key, wantValue, gotValue)
				}
			}
		})
	}
}

// TestProcess テンプレートの処理をテストします。
func TestProcess(t *testing.T) {
	// テストケース
	testCases := []struct {
		name      string
		content   string
		variables map[string]interface{}
		want      string
		wantErr   bool
	}{}

	// 正常系：単純な変数置換
	testCases = append(testCases, struct {
		name      string
		content   string
		variables map[string]interface{}
		want      string
		wantErr   bool
	}{
		name:    "simple variable replacement",
		content: "Project: {{.ProjectName}}, Module: {{.ModuleName}}",
		variables: map[string]interface{}{
			"ProjectName": "test-project",
			"ModuleName":  "github.com/user/test",
		},
		want:    "Project: test-project, Module: github.com/user/test",
		wantErr: false,
	})

	// 正常系：複数の変数
	testCases = append(testCases, struct {
		name      string
		content   string
		variables map[string]interface{}
		want      string
		wantErr   bool
	}{
		name:    "multiple variables",
		content: "{{.ProjectName}} is a {{.ProjectType}} project using {{.Framework}}",
		variables: map[string]interface{}{
			"ProjectName": "my-api",
			"ProjectType": "api",
			"Framework":   "gin",
		},
		want:    "my-api is a api project using gin",
		wantErr: false,
	})

	// 異常系：無効なテンプレート構文
	testCases = append(testCases, struct {
		name      string
		content   string
		variables map[string]interface{}
		want      string
		wantErr   bool
	}{
		name:    "invalid template syntax",
		content: "{{.ProjectName",
		variables: map[string]interface{}{
			"ProjectName": "tes",
		},
		want:    "",
		wantErr: true,
	})

	// 異常系：変数が存在しない場合
	testCases = append(testCases, struct {
		name      string
		content   string
		variables map[string]interface{}
		want      string
		wantErr   bool
	}{
		name:    "missing variable",
		content: "{{.ProjectName}} {{.Missing}}",
		variables: map[string]interface{}{
			"ProjectName": "test",
		},
		want:    "test <no value>", // text/templateのデフォルト動作
		wantErr: false,
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processor := NewTemplateProcessor()
			got, err := processor.Process(tc.content, tc.variables)
			if (err != nil) != tc.wantErr {
				t.Fatalf("期待されるエラー状態と異なります。wantErr: %v, gotErr: %v", tc.wantErr, err != nil)
			}
			if got != tc.want {
				t.Errorf("処理結果が期待値と異なります。期待値: %s, 実際の値: %s", tc.want, got)
			}
		})
	}
}

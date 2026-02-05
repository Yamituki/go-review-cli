package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Yamituki/go-review-cli/internal/domain/value"
)

type ProjectInput struct {
	Name        string
	Type        string
	Framework   string
	Module      string
	Description string
	Path        string
}

// PromptProjectCreation 対話形式でプロジェクト作成に必要な情報を入力させる
func PromptProjectCreation() (*ProjectInput, error) {
	// プロジェクト情報入力会話のセットアップ
	questions := []*survey.Question{
		{
			Name: "projectName",
			Prompt: &survey.Input{
				Message: "プロジェクト名を入力してください:",
			},
			Validate: survey.Required,
		},
		{
			Name: "Module",
			Prompt: &survey.Input{
				Message: "モジュールパスを入力してください (例: github.com/username/project):",
			},
			Validate: survey.Required,
		},
		{
			Name: "Description",
			Prompt: &survey.Input{
				Message: "プロジェクトの説明を入力してください:",
			},
		},
		{
			Name: "Path",
			Prompt: &survey.Input{
				Message: "プロジェクトの保存先パスを入力してください:",
				Default: ".",
			},
		},
	}

	// ユーザーからの入力を格納する変数
	answers := struct {
		ProjectName string
		Module      string
		Description string
		Path        string
	}{}

	// プロジェクト情報入力会話を実行
	err := survey.Ask(questions, &answers)
	if err != nil {
		return nil, err
	}

	// プロジェクトタイプ選択会話のセットアップ
	typeQuestion := &survey.Select{
		Message: "プロジェクトのタイプを選択してください:",
		Options: []string{
			value.ProjectTypeAPI.String(),
			value.ProjectTypeCLI.String(),
			value.ProjectTypeMicroservice.String(),
		},
	}

	// プロジェクトタイプの選択を格納する変数
	var projectType string

	// プロジェクトタイプ選択会話を実行
	err = survey.AskOne(typeQuestion, &projectType)
	if err != nil {
		return nil, err
	}

	// フレームワークの選択を格納する変数
	var frameworkType string

	// フレームワーク選択会話のセットアップ
	if projectType == value.ProjectTypeAPI.String() {
		// フレームワーク選択会話のセットアップ
		frameworkQuestion := &survey.Select{
			Message: "使用するフレームワークを選択してください:",
			Options: []string{
				value.FrameworkTypeGin.String(),
				value.FrameworkTypeEcho.String(),
			},
		}

		// フレームワーク選択会話を実行
		err = survey.AskOne(frameworkQuestion, &frameworkType)
		if err != nil {
			return nil, err
		}
	}

	// 入力されたプロジェクト情報を返す
	input := &ProjectInput{
		Name:        answers.ProjectName,
		Type:        projectType,
		Framework:   frameworkType,
		Module:      answers.Module,
		Description: answers.Description,
		Path:        answers.Path,
	}

	return input, nil
}

package prompts

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// PromptTemplateRemove テンプレート削除のプロンプト
func PromptTemplateRemove(templateName string) (bool, error) {
	// 確認
	confirm := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("テンプレート '%s' を削除するか？", templateName),
		Default: false,
	}

	err := survey.AskOne(prompt, &confirm)
	if err != nil {
		return false, err
	}

	return confirm, nil
}

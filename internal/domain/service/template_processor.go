package service

import (
	"strings"

	"github.com/Yamituki/go-review-cli/internal/domain/entity"
)

type TemplateProcessor struct{}

// NewTemplateProcessor 新たなTemplateProcessorを生成する
func NewTemplateProcessor() *TemplateProcessor {
	return &TemplateProcessor{}
}

// PrepareVariables テンプレート変数の準備を行う
func (p *TemplateProcessor) PrepareVariables(project *entity.Project) map[string]interface{} {
	// 作成
	variables := make(map[string]interface{})

	// 内容
	variables["ProjectName"] = project.Name
	variables["ModuleName"] = project.Module
	variables["ProjectType"] = project.Type
	variables["Framework"] = project.Framework
	variables["Description"] = project.Description

	return variables
}

// Process テンプレートの処理を行う
func (p *TemplateProcessor) Process(content string, variables map[string]interface{}) (string, error) {
	// text/templateなどを使用してテンプレート処理を実装する
	// {{.変数名}} の形式で変数を置換する簡易実装例

	result := content
	for key, value := range variables {
		placeholder := "{{." + key + "}}"
		result = strings.ReplaceAll(result, placeholder, value.(string))
	}
	return result, nil
}

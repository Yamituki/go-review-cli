package service

import (
	"bytes"
	"text/template"

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
	variables["ProjectName"] = project.Name.String()
	variables["ModuleName"] = project.Module.String()
	variables["ProjectType"] = project.Type.String()
	variables["Description"] = project.Description

	if project.Framework != nil {
		variables["Framework"] = project.Framework.String()
	} else {
		variables["Framework"] = ""
	}

	return variables
}

// Process テンプレートの処理を行う
func (p *TemplateProcessor) Process(content string, variables map[string]interface{}) (string, error) {
	temp, err := template.New("tmpl").Parse(content)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	err = temp.Execute(&buffer, variables)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

package entity

import (
	"fmt"

	"github.com/Yamituki/go-review-cli/internal/domain/value"
)

type Project struct {
	Name        value.ProjectName
	Type        value.ProjectType
	Framework   *value.FrameworkType
	Module      value.ModuleName
	Description string
	Path        string
}

// NewProject　新なプロジェクトエンティティを生成する
func NewProject(
	name string,
	projectType string,
	framework string,
	module string,
	description string,
	path string,
) (*Project, error) {

	projectName, err := value.NewProjectName(name)
	if err != nil {
		return nil, err
	}

	projectTypeValue, err := value.NewProjectType(projectType)
	if err != nil {
		return nil, err
	}

	var frameworkValue *value.FrameworkType
	if framework != "" {
		fw, err := value.NewFrameworkType(framework)
		if err != nil {
			return nil, err
		}
		frameworkValue = &fw
	}

	moduleName, err := value.NewModuleName(module)
	if err != nil {
		return nil, err
	}

	p := &Project{
		Name:        projectName,
		Type:        projectTypeValue,
		Framework:   frameworkValue,
		Module:      moduleName,
		Description: description,
		Path:        path,
	}

	// Typeがapiの場合、Frameworkは必須
	if projectTypeValue == value.ProjectTypeAPI && frameworkValue == nil {
		return nil, fmt.Errorf("フレームワークは必須です")
	} else if projectTypeValue != value.ProjectTypeAPI && frameworkValue != nil {
		return nil, fmt.Errorf("フレームワークは不要です")
	}

	return p, nil
}

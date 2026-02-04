package value

import "fmt"

type ProjectType string

const (
	ProjectTypeAPI          ProjectType = "api"
	ProjectTypeCLI          ProjectType = "cli"
	ProjectTypeMicroservice ProjectType = "microservice"
)

// NewProjectType 新しいプロジェクトタイプを生成します。
func NewProjectType(t string) (ProjectType, error) {
	pt := ProjectType(t)

	if !pt.IsValid() {
		return "", fmt.Errorf("無効なプロジェクトタイプです: %s", t)
	}

	return pt, nil
}

// String プロジェクトタイプを文字列として返します。
func (p ProjectType) String() string {
	return string(p)
}

// IsValid プロジェクトタイプが有効かどうかを確認します。
func (p ProjectType) IsValid() bool {
	switch p {
	case ProjectTypeAPI, ProjectTypeCLI, ProjectTypeMicroservice:
		return true
	default:
		return false
	}
}

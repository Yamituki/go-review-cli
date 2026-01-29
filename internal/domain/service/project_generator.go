package service

import "github.com/Yamituki/go-review-cli/internal/domain/entity"

type ProjectGenerator struct{}

// NewProjectGenerator 新たなProjectGeneratorを生成する
func NewProjectGenerator() *ProjectGenerator {
	return &ProjectGenerator{}
}

// ValidateProject プロジェクトのバリデーションを行う
func (g *ProjectGenerator) ValidateProject(project *entity.Project) error {
	return project.Validate()
}

// GenerateStructure プロジェクトのディレクトリ構造を生成する
func (g *ProjectGenerator) GenerateStructure(project *entity.Project) (map[string]string, error) {
	// ここにディレクトリ構造生成のロジックを実装する
	structure := make(map[string]string)

	// ルーツディレクトリ
	structure["root"] = project.Path + "/" + string(project.Name)

	// 共通ディレクトリ
	structure["cmd"] = structure["root"] + "/cmd"
	structure["internal"] = structure["root"] + "/internal"
	structure["pkg"] = structure["root"] + "/pkg"

	return structure, nil
}

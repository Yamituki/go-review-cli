package service

import (
	"fmt"

	"github.com/Yamituki/go-review-cli/internal/domain/entity"
)

type ProjectGenerator struct{}

// NewProjectGenerator 新たなProjectGeneratorを生成する
func NewProjectGenerator() *ProjectGenerator {
	return &ProjectGenerator{}
}

// GenerateStructure プロジェクトのディレクトリ構造を生成する
func (g *ProjectGenerator) GenerateStructure(project *entity.Project) (map[string]string, error) {
	//プロジェクトルート/
	// ├── cmd/
	// ├── internal/
	// │   ├── domain/
	// │   ├── usecase/
	// │   ├── handler/
	// │   └── infrastructure/
	// ├── pkg/
	// ├── test/
	// ├── configs/
	// ├── scripts/
	// └── docs/

	// ここにディレクトリ構造生成のロジックを実装する
	structure := make(map[string]string)

	// ルーツディレクトリ
	structure["root"] = fmt.Sprintf("%s/%s", project.Path, project.Name.String())

	// 共通ディレクトリ
	structure["cmd"] = fmt.Sprintf("%s/cmd", structure["root"])
	structure["internal"] = fmt.Sprintf("%s/internal", structure["root"])
	structure["pkg"] = fmt.Sprintf("%s/pkg", structure["root"])
	structure["test"] = fmt.Sprintf("%s/test", structure["root"])
	structure["configs"] = fmt.Sprintf("%s/configs", structure["root"])
	structure["scripts"] = fmt.Sprintf("%s/scripts", structure["root"])
	structure["docs"] = fmt.Sprintf("%s/docs", structure["root"])

	// internal以下のディレクトリ
	structure["internal/domain"] = fmt.Sprintf("%s/domain", structure["internal"])
	structure["internal/usecase"] = fmt.Sprintf("%s/usecase", structure["internal"])
	structure["internal/handler"] = fmt.Sprintf("%s/handler", structure["internal"])
	structure["internal/infrastructure"] = fmt.Sprintf("%s/infrastructure", structure["internal"])

	return structure, nil
}

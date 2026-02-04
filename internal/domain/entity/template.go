package entity

import "fmt"

type Template struct {
	Name        string
	Version     string
	Description string
	Language    string
	Type        string
	Path        string
}

// NewTemplate 新たなTemplateエンティティを生成
func NewTemplate(name, version, description, language, templateType, path string) (*Template, error) {
	t := &Template{
		Name:        name,
		Version:     version,
		Description: description,
		Language:    language,
		Type:        templateType,
		Path:        path,
	}

	if err := t.Validate(); err != nil {
		return nil, err
	}

	return t, nil
}

// Validate テンプレートのバリデーションを行う
func (t *Template) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("templateのNameが空です")
	}

	if t.Version == "" {
		return fmt.Errorf("templateのVersionが空です")
	}

	return nil
}

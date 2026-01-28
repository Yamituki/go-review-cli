package value

import "fmt"

type ProjectName string

// NewProjectName 新しいプロジェクト名を生成します。
func NewProjectName(name string) (ProjectName, error) {
	pn := ProjectName(name)

	if err := pn.Validate(); err != nil {
		return "", err
	}

	return pn, nil
}

// String プロジェクト名を文字列として返します。
func (p ProjectName) String() string {
	return string(p)
}

// Validate プロジェクト名のバリデーションを行います。
func (p ProjectName) Validate() error {
	// 一文字以上であることを確認
	if len(p) == 0 {
		return fmt.Errorf("プロジェクト名は1文字以上である必要があります")
	}

	// 英数字とハイフンのみを許可
	for _, char := range p {

		// 英字
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			continue
		}

		// 数字
		if char >= '0' && char <= '9' {
			continue
		}

		// ハイフン
		if char == '-' {
			continue
		}

		return fmt.Errorf("プロジェクト名には英数字とハイフンのみを使用できます")
	}

	// 先頭にハイフンがないことを確認
	if p[0] == '-' {
		return fmt.Errorf("プロジェクト名の先頭にハイフンを使用できません")
	}

	// 末尾にハイフンがないことを確認
	if p[len(p)-1] == '-' {
		return fmt.Errorf("プロジェクト名の末尾にハイフンを使用できません")
	}

	return nil
}

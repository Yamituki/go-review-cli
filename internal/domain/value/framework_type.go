package value

import "fmt"

type FrameworkType string

const (
	FrameworkTypeGin  FrameworkType = "gin"
	FrameworkTypeEcho FrameworkType = "echo"
)

// NewFrameworkType 新しいフレームワークタイプを生成します。
func NewFrameworkType(t string) (FrameworkType, error) {
	ft := FrameworkType(t)

	if !ft.IsValid() {
		return "", fmt.Errorf("無効なフレームワークタイプです: %s", t)
	}

	return ft, nil
}

// String フレームワークタイプを文字列として返します。
func (f FrameworkType) String() string {
	return string(f)
}

// IsValid フレームワークタイプが有効かどうかを確認します。
func (f FrameworkType) IsValid() bool {
	switch f {
	case FrameworkTypeGin, FrameworkTypeEcho:
		return true
	default:
		return false
	}
}

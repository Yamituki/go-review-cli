package value

import (
	"fmt"
	"strings"
)

type ModuleName string

// NewModuleName 新しいモジュール名を生成します。
func NewModuleName(name string) (ModuleName, error) {
	mn := ModuleName(name)

	if err := mn.Validate(); err != nil {
		return "", err
	}

	return mn, nil
}

// String モジュール名を文字列として返します。
func (m ModuleName) String() string {
	return string(m)
}

// Validate モジュール名のバリデーションを行います。
func (m ModuleName) Validate() error {
	// Go Module形式かどうかを簡易的にチェック
	// 例: github.com/user/repo

	// スラッシュで区切られた３つ以上のセグメントをチェック
	segments := strings.Split(string(m), "/")
	if len(segments) < 3 {
		return fmt.Errorf("モジュール名は少なくとも3つのセグメントを含む必要があります (例: github.com/user/repo)")
	}

	return nil
}

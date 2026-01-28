package value

import "testing"

// TestNewModuleName 新しいモジュール名を生成するテスト
func TestNewModuleName(t *testing.T) {
	// テストケース
	testCases := []string{
		"github.com/user/repo", // 正常系
		"",                     // 空文字列
		"github.com/user",      // スラッシュが3つ未満
		"github.com-user-repo", // 不正な形式
	}

	for i, tc := range testCases {
		_, err := NewModuleName(tc)

		// 正常系
		if i == 0 && err != nil {
			t.Errorf("テストケース %d: 期待されるエラーなし、実際のエラー: %v", i, err)
		}

		// 異常系
		if i != 0 && err == nil {
			t.Errorf("テストケース %d: 期待されるエラーあり、実際のエラーなし", i)
		}
	}
}

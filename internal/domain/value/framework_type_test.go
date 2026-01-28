package value

import "testing"

// TestNewFrameworkType 新しい FrameworkType を作成するテスト
func TestNewFrameworkType(t *testing.T) {
	// テストケース
	testCases := []string{
		// 正常系
		"gin",
		"echo",
		// 異常系
		"",
		"Flask",
		"Django",
	}

	for i, tc := range testCases {
		_, err := NewFrameworkType(tc)

		// 正常系
		if i < 2 && err != nil {
			t.Errorf("テストケース %d: 期待されるエラーなし、実際のエラー: %v", i, err)
		}

		// 異常系
		if i >= 2 && err == nil {
			t.Errorf("テストケース %d: 期待されるエラーあり、実際のエラーなし", i)
		}

	}
}

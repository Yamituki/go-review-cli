package value

import "testing"

// TestNewProjectType 新しい ProjectType を作成するテスト
func TestNewProjectType(t *testing.T) {
	// テストケース
	testCases := []string{
		// 正常系
		"api",
		"cli",
		"microservice",
		// 異常系
		"",
		"webapp",
	}

	for i, tc := range testCases {
		_, err := NewProjectType(tc)

		// 正常系
		if i < 3 && err != nil {
			t.Errorf("テストケース %d: 期待されるエラーなし、実際のエラー: %v", i, err)
		}

		// 異常系
		if i >= 3 && err == nil {
			t.Errorf("テストケース %d: 期待されるエラーあり、実際のエラーなし", i)
		}
	}

}

package value

import "testing"

// TestNewProjectName 新しいプロジェクト名を生成するテスト用ヘルパー関数。
func TestNewProjectName(t *testing.T) {
	// テストケース
	testCases := []string{
		"my-test-project",  // 正常系
		"",                 // 空文字列
		"my_test_project!", // 無効な文字列（特殊文字を含む）
		"-mytestproject",   // 先頭にハイフンを含む
		"mytestproject-",   // 末尾にハイフンを含む
	}

	for i, tc := range testCases {
		_, err := NewProjectName(tc)

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

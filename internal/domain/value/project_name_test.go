package value

import "testing"

// TestNewProjectName 新しいプロジェクト名を生成するテスト
func TestNewProjectName(t *testing.T) {
	// テストケース
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"正常: 有効なプロジェクト名", "my-test-project", false},
		{"異常: 空文字列", "", true},
		{"異常: 特殊文字", "my_test_project!", true},
		{"異常: 先頭ハイフン", "-mytestproject", true},
		{"異常: 末尾ハイフン", "mytestproject-", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProjectName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("期待: エラー=%v, 実際: %v", tt.wantErr, err)
			}
		})
	}
}

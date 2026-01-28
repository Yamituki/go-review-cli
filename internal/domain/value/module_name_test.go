package value

import "testing"

// TestNewModuleName 新しいモジュール名を生成するテスト
func TestNewModuleName(t *testing.T) {
	// テストケース
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"正常: 有効なモジュール名", "github.com/user/repo", false},
		{"異常: 空文字列", "", true},
		{"異常: スラッシュが3つ未満", "github.com/user", true},
		{"異常: 不正な形式", "github.com-user-repo", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewModuleName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("期待: エラー=%v, 実際: %v", tt.wantErr, err)
			}
		})
	}
}

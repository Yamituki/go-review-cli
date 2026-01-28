package value

import "testing"

// TestNewFrameworkType 新しい FrameworkType を作成するテスト
func TestNewFrameworkType(t *testing.T) {
	// テストケース
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"正常: 有効なフレームワーク 'gin'", "gin", false},
		{"正常: 有効なフレームワーク 'echo'", "echo", false},
		{"異常: 空文字列", "", true},
		{"異常: 無効なフレームワーク 'Flask'", "Flask", true},
		{"異常: 無効なフレームワーク 'Django'", "Django", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewFrameworkType(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("期待: エラー=%v, 実際: %v", tt.wantErr, err)
			}
		})
	}
}

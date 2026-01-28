package value

import "testing"

// TestNewProjectType 新しい ProjectType を作成するテスト
func TestNewProjectType(t *testing.T) {
	// テストケース
	testCases := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"正常: 有効なプロジェクトタイプ 'api'", "api", false},
		{"正常: 有効なプロジェクトタイプ 'cli'", "cli", false},
		{"正常: 有効なプロジェクトタイプ 'microservice'", "microservice", false},
		{"異常: 空文字列", "", true},
		{"異常: 無効なプロジェクトタイプ 'webapp'", "webapp", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProjectType(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("期待: エラー=%v, 実際: %v", tt.wantErr, err)
			}
		})
	}

}

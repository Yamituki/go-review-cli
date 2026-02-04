package repository

type ConfigRepository interface {
	// Get 設定値を取得する
	Get(key string) (string, error)
	// Set 設定値を保存する
	Set(key string, value string) error
	// GetAll すべての設定値を取得する
	GetAll() (map[string]string, error)
}

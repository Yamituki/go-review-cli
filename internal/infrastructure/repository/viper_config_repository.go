package repository

import (
	"fmt"

	"github.com/spf13/viper"
)

type ViperConfigRepository struct{}

// NewViperConfigRepository 新たなViperConfigRepositoryを生成します。
func NewViperConfigRepository() *ViperConfigRepository {
	return &ViperConfigRepository{}
}

// Get 設定値を取得します。
func (r *ViperConfigRepository) Get(key string) (string, error) {
	value := viper.GetString(key)

	if value == "" {
		return "", fmt.Errorf("設定が見つかりません: %s", key)
	}

	return value, nil
}

// Set 設定値を設定します。
func (r *ViperConfigRepository) Set(key string, value string) error {
	viper.Set(key, value)

	return viper.WriteConfig()
}

// GetAll 全ての設定値を取得します。
func (r *ViperConfigRepository) GetAll() (map[string]string, error) {
	settings := viper.AllSettings()
	result := make(map[string]string)
	for key, value := range settings {
		strValue, ok := value.(string)
		if !ok {
			continue
		}
		result[key] = strValue
	}
	return result, nil
}

package repository

import (
	"fmt"
	"os"
	"path/filepath"

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

// EnsureConfigFile 設定ファイルが存在しない場合は作成します。
func (r *ViperConfigRepository) EnsureConfigFile() error {
	err := viper.ReadInConfig()
	if err == nil {
		return nil
	}

	root, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(root, ".go-review-cli")

	// ディレクトリが存在しない場合は作成します。
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	configFilePath := filepath.Join(configPath, "config.yaml")

	// ファイルが既に存在するかチェック
	if _, err := os.Stat(configFilePath); err == nil {
		return nil
	}

	viper.SetDefault("project.default_path", ".")
	viper.SetDefault("project.default_framework", "gin")

	return viper.SafeWriteConfigAs(configFilePath)

}

// Reset 設定ファイルをリセットします。
func (r *ViperConfigRepository) Reset() error {
	viper.Set("project.default_path", ".")
	viper.Set("project.default_framework", "gin")

	return viper.WriteConfig()
}

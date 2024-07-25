package model

import (
	"cell_phone_store/structure"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

var AppConfigTable = "app_config"

func GetBaseConfig() (structure.BaseConfig, error) {
	var baseConfigRecord struct {
		ID          int
		ConfigValue string `gorm:"column:config_value;type:jsonb"`
		CreatedAt   string
		UpdatedAt   string
	}
	result := DB.Table(AppConfigTable).Where("config_key = ?", "base_config").First(&baseConfigRecord)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return structure.BaseConfig{}, fmt.Errorf("无记录: %w", result.Error)
		}
		return structure.BaseConfig{}, result.Error
	}
	var configValue structure.ConfigValue
	if err := json.Unmarshal([]byte(baseConfigRecord.ConfigValue), &configValue); err != nil {
		return structure.BaseConfig{}, err
	}
	return structure.BaseConfig{
		ID:          baseConfigRecord.ID,
		ConfigValue: configValue,
		CreatedAt:   baseConfigRecord.CreatedAt,
		UpdatedAt:   baseConfigRecord.UpdatedAt,
	}, nil
}
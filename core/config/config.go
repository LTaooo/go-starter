package config

import (
	"go-starter/core/logger"

	"github.com/spf13/viper"
)

var configData *AppConfig

func GetConfig() *AppConfig {
	return configData
}

func LoadConfig() {
	// 1. 配置 Viper 读取 YAML 文件
	viper.SetConfigName("config") // 配置文件名 (不带扩展名)
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 搜索路径

	// 2. 设置环境变量自动绑定和替换
	viper.AutomaticEnv()

	// 3. 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		logger.SugaredLogger.Error("读取配置文件失败", "error", err)
		panic(err)
	}

	// 4. 解析成结构体
	cfg := NewAppConfig()
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.SugaredLogger.Error("解析配置失败", "error", err)
		panic(err)
	}

	configData = cfg
}

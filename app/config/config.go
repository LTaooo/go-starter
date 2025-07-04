package config

import (
	"go-starter/core/logger"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var configData *AppConfig

func GetConfig() *AppConfig {
	return configData
}

func LoadConfig() {
	// 1. 加载 .env 文件
	if err := godotenv.Load(".env"); err != nil {
		logger.SugaredLogger.Error("读取 .env 文件失败", "error", err)
	}

	// 2. 配置 Viper
	viper.SetConfigName(".env") // 配置文件名 (不带扩展名)
	viper.SetConfigType("env")  // 配置文件类型
	viper.AddConfigPath(".")    // 搜索路径

	// 3. 设置环境变量自动绑定和替换
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 5. 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		logger.SugaredLogger.Error("读取配置文件失败", "error", err)
	}

	// 6. 解析成结构体
	cfg := NewAppConfig()
	if err := viper.Unmarshal(&cfg); err != nil {
		logger.SugaredLogger.Error("解析配置失败", "error", err)
	}
	configData = cfg
}

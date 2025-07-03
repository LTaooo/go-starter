package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

var GlobalConfig *AppConfig

func LoadConfig() {
	// 1. 加载 .env 文件（或系统环境变量）
	_ = godotenv.Load(".env")
	viper.AutomaticEnv()

	// 3. 创建 nacos 客户端
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(viper.GetString("NACOS_IP"), uint64(viper.GetInt("NACOS_PORT"))),
	}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(viper.GetString("NACOS_NAMESPACE")),
		constant.WithUsername(viper.GetString("NACOS_USERNAME")),
		constant.WithPassword(viper.GetString("NACOS_PASSWORD")),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("logs"),
		constant.WithCacheDir("runtime"),
		constant.WithLogLevel("debug"),
	)

	nacosClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)
	if err != nil {
		log.Fatalf("Create nacos client failed: %v", err)
	}

	// 4. 读取远程配置（yaml）
	content, err := nacosClient.GetConfig(vo.ConfigParam{
		DataId: viper.GetString("NACOS_DATA_ID"),
		Group:  viper.GetString("NACOS_GROUP"),
	})
	if err != nil {
		log.Fatalf("Get config from nacos failed: %v", err)
	}

	// 5. 加载到 viper
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(content)); err != nil {
		log.Fatalf("Read nacos config failed: %v", err)
	}

	// 6. 设置默认值（可选）
	viper.SetDefault("server.port", 3000)

	// 7. 解析成结构体
	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unmarshal config failed: %v", err)
	}
	GlobalConfig = &cfg

	fmt.Printf("✅ Config loaded: %+v\n", cfg)
}

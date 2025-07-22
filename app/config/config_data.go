package config

type MysqlConfig struct {
	Enable      bool   `mapstructure:"enable"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Database    string `mapstructure:"database"`
	PoolSize    int    `mapstructure:"pool_size"`
	MaxPoolSize int    `mapstructure:"max_pool_size"`
}

func newMysqlConfig() *MysqlConfig {
	return &MysqlConfig{
		Enable:      false,
		Host:        "localhost",
		Port:        3306,
		Username:    "root",
		Password:    "root",
		Database:    "go_starter",
		PoolSize:    10,
		MaxPoolSize: 100,
	}
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

func newRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		Database: 0,
	}
}

type AppConfig struct {
	AppName string      `mapstructure:"app_name"`
	AppEnv  string      `mapstructure:"app_env"`
	AppPort string      `mapstructure:"app_port"`
	AppHost string      `mapstructure:"app_host"`
	Mysql   MysqlConfig `mapstructure:"mysql"`
	Redis   RedisConfig `mapstructure:"redis"`
}

func (c *AppConfig) GetListenAddr() string {
	return c.AppHost + ":" + c.AppPort
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		AppName: "go-starter",
		AppEnv:  "dev",
		AppPort: "8000",
		AppHost: "0.0.0.0",
		Mysql:   *newMysqlConfig(),
		Redis:   *newRedisConfig(),
	}
}

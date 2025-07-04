package config

type AppConfig struct {
	AppName string `mapstructure:"app_name"`
	AppEnv  string `mapstructure:"app_env"`
	AppPort string `mapstructure:"app_port"`
	AppHost string `mapstructure:"app_host"`
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
	}
}

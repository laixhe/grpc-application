package config

// 替换环境变量

// AppConfig 项目配置
const (
	EnvAppVersion = "APP_VERSION"
)

// EnvConfig 环境变量和配置文件对应
type EnvConfig struct {
	env       string // 环境变量
	configKey string // 配置文件对应
}

// ListEnvConfig 环境变量和配置文件对应列表
func ListEnvConfig() []EnvConfig {
	return []EnvConfig{
		// AppConfig 项目配置
		{env: EnvAppVersion, configKey: "app.version"},
		//
	}
}

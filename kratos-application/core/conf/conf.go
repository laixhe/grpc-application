package conf

import (
	kratosConfig "github.com/go-kratos/kratos/v2/config"
	kratosConfigFile "github.com/go-kratos/kratos/v2/config/file"

	"kratos-application/api/gen/config"
	"kratos-application/api/gen/config/capp"
	"kratos-application/api/gen/config/cauth"
	"kratos-application/api/gen/config/cdb"
	"kratos-application/api/gen/config/clog"
	"kratos-application/api/gen/config/cmongodb"
	"kratos-application/api/gen/config/credis"
	"kratos-application/api/gen/config/cserver"
)

var cc *config.Config

// Get Config
func Get() *config.Config {
	return cc
}

// NewFileConfigSource 创建一个本地文件配置源
func NewFileConfigSource(filePath string) kratosConfig.Source {
	return kratosConfigFile.NewSource(filePath)
}

// Init 初始化配置
func Init(configPath string) {
	c := kratosConfig.New(
		kratosConfig.WithSource(
			NewFileConfigSource(configPath),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	cc = &config.Config{
		App:     &capp.App{},
		Log:     &clog.Log{},
		Servers: &cserver.Servers{},
		Jwt:     &cauth.Jwt{},
		Db:      &cdb.DB{},
		Mongodb: &cmongodb.MongoDB{},
		Redis:   &credis.Redis{},
	}

	if err := c.Scan(cc); err != nil {
		panic(err)
	}

	Checking()
}

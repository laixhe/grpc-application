package core

import (
	"errors"
	"fmt"

	kratosConfig "github.com/go-kratos/kratos/v2/config"
	kratosConfigFile "github.com/go-kratos/kratos/v2/config/file"

	"github.com/laixhe/gonet/jwt"
	"github.com/laixhe/gonet/orm/orm"
)

type Addr struct {
	IP      string `json:"ip"`
	Port    int    `json:"port"`
	Timeout int    `json:"timeout"`
}

func (a *Addr) Addr() string {
	return fmt.Sprintf("%s:%d", a.IP, a.Port)
}

type Config struct {
	Jwt  *jwt.Config `json:"jwt"`
	Orm  *orm.Config `json:"orm"`
	Http *Addr       `json:"http"`
	Grpc *Addr       `json:"grpc"`
}

func (c *Config) Check() error {
	if err := c.Jwt.Check(); err != nil {
		return err
	}
	if err := c.Orm.Check(); err != nil {
		return err
	}
	if c.Http == nil {
		return errors.New("http config is nil")
	}
	if c.Http.Port <= 0 {
		return errors.New("http port is invalid")
	}
	if c.Grpc == nil {
		return errors.New("grpc config is nil")
	}
	if c.Grpc.Port <= 0 {
		return errors.New("grpc port is invalid")
	}
	return nil
}

func NewConfig(configPath string) *Config {
	config := kratosConfig.New(kratosConfig.WithSource(kratosConfigFile.NewSource(configPath)))
	if err := config.Load(); err != nil {
		panic(err)
	}
	c := &Config{}
	if err := config.Scan(c); err != nil {
		panic(err)
	}
	if err := c.Check(); err != nil {
		panic(err)
	}
	return c
}

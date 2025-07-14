package core

import (
	"github.com/laixhe/gonet/orm/mysql"
	"github.com/laixhe/gonet/orm/orm"
)

// DEFAULT 默认key
const DEFAULT = "default"

// Server 服务
type Server struct {
	config *Config
	orm    map[string]orm.Client
}

// Config 获取配置
func (s *Server) Config() *Config {
	return s.config
}

// 初始化ORM
func (s *Server) initOrm(config *orm.Config, key ...string) error {
	db, err := mysql.Init(config, &OrmWriter{}, "")
	if err != nil {
		return err
	}
	if len(key) > 0 {
		s.orm[key[0]] = db
	} else {
		s.orm[DEFAULT] = db
	}
	return nil
}

// Orm 获取ORM
func (s *Server) Orm(key ...string) orm.Client {
	if len(key) > 0 {
		return s.orm[key[0]]
	} else {
		return s.orm[DEFAULT]
	}
}

func (s *Server) init() *Server {
	if err := s.initOrm(s.config.Orm); err != nil {
		panic(err)
	}
	return s
}

// NewServer 创建服务
func NewServer(configFile string) *Server {
	s := &Server{
		config: NewConfig(configFile),
		orm:    make(map[string]orm.Client),
	}
	return s.init()
}

package redisx

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	redis "github.com/redis/go-redis/v9"
)

type Config struct {
	Addr        string `mapstructure:"addr"`
	DbNum       uint   `mapstructure:"db_num"`        // 选择N号数据库
	Password    string `mapstructure:"password"`      // 密码
	PoolSize    uint   `mapstructure:"pool_size"`     // 最大链接数
	MinIdleConn uint   `mapstructure:"min_idle_conn"` // 空闲链接数
}

// Redisx 客户端
type Redisx struct {
	client *redis.Client
}

// Ping 判断服务是否可用
func (r *Redisx) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := r.client.Ping(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}

// Client get redis client
func (r *Redisx) Client() *redis.Client {
	return r.client
}

var db *Redisx

// Client get redis client
func Client() *redis.Client {
	return db.client
}

// Ping 判断服务是否可用
func Ping() error {
	return db.Ping()
}

// initSingle 单机
func initSingle(c *Config) *redis.Client {
	options := &redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       int(c.DbNum),
	}
	if c.PoolSize > 0 {
		options.PoolSize = int(c.PoolSize)
	}
	if c.MinIdleConn > 0 {
		options.MinIdleConns = int(c.MinIdleConn)
	}
	return redis.NewClient(options)
}

// Connect 连接数据库
func Connect(c *Config) (*Redisx, error) {
	r := &Redisx{}
	r.client = initSingle(c)
	err := r.Ping()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func Init(c *Config) {
	log.Debugf("redis Config=%v", c)
	log.Debug("redis 开始初始化...")

	var err error
	db, err = Connect(c)
	if err != nil {
		panic(err)
	}

	log.Debug("redis 初始化完毕...")
}

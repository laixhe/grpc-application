package redisx

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	redis "github.com/redis/go-redis/v9"

	"kratos-application/api/gen/config/credis"
)

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
func initSingle(c *credis.Redis) *redis.Client {
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
func Connect(c *credis.Redis) (*Redisx, error) {
	r := &Redisx{}
	r.client = initSingle(c)
	err := r.Ping()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func Init(c *credis.Redis) {
	if c == nil {
		panic(errors.New("config redis is nil"))
	}
	if c.Addr == "" {
		panic(errors.New("config redis addr is nil"))
	}
	log.Debugf("redis Config=%v", c)
	log.Debug("redis 开始初始化...")

	var err error
	db, err = Connect(c)
	if err != nil {
		panic(err)
	}

	log.Debug("redis 初始化完毕...")
}

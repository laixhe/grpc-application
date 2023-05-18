package gormx

import (
	stdLog "log"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Config struct {
	Dsn          string `mapstructure:"dsn"`            // 连接地址
	MaxIdleCount uint   `mapstructure:"max_idle_count"` // 设置空闲连接池中连接的最大数量
	MaxOpenCount uint   `mapstructure:"max_open_count"` // 设置打开数据库连接的最大数量
	MaxLifeTime  uint   `mapstructure:"max_life_time"`  // 设置了连接可复用的最大时间(要比数据库设置连接超时时间少)(单位秒)
}

// Gormx 客户端
type Gormx struct {
	client *gorm.DB
}

// Ping 判断服务是否可用
func (g *Gormx) Ping() error {
	sqlDB, err := g.client.DB()
	if err != nil {
		return err
	}
	// 验证数据库连接是否正常
	return sqlDB.Ping()
}

var db *Gormx

// Client get db client
func Client() *gorm.DB {
	return db.client
}

// Connect 连接数据库
func Connect(c *Config) (*Gormx, error) {

	defaultLogger := logger.New(NewWriter(stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	client, err := gorm.Open(mysql.Open(c.Dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: defaultLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := client.DB()
	if err != nil {
		return nil, err
	}
	if c.MaxIdleCount > 0 {
		// 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(int(c.MaxIdleCount))
	}
	if c.MaxOpenCount > 0 {
		// 设置打开数据库连接的最大数量
		sqlDB.SetMaxOpenConns(int(c.MaxOpenCount))
	}
	if c.MaxLifeTime > 0 {
		// 设置了连接可复用的最大时间(要比数据库设置连接超时时间少)
		sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifeTime) * time.Second)
	}
	// 验证数据库连接是否正常
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	return &Gormx{
		client: client,
	}, nil
}

// Init 初始化数据库
func Init(c *Config) {
	log.Debugf("db Config=%v", c)
	log.Debug("db 开始初始化...")

	var err error
	db, err = Connect(c)
	if err != nil {
		panic(err)
	}

	log.Debug("db 初始化完毕...")
}

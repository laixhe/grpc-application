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

// Client get db client
func (g *Gormx) Client() *gorm.DB {
	return g.client
}

// Table 定义表名
func (g *Gormx) Table(name string, args ...interface{}) *gorm.DB {
	return g.client.Table(name, args...)
}

// Model 指定操作模型
func (g *Gormx) Model(value interface{}) *gorm.DB {
	return g.client.Model(value)
}

// Create 创建数据
func (g *Gormx) Create(value interface{}) error {
	return g.client.Create(value).Error
}

// First 获取第一条记录（主键升序），且没有找到记录时，它会返回 gorm.ErrRecordNotFound 错误
func (g *Gormx) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return g.client.First(dest, conds...)
}

// Find 获取多条记录（查找所有符合给定条件的记录）
func (g *Gormx) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return g.client.Find(dest, conds...)
}

// Select 选择指定查询、创建、更新时所需的字段，默认情况下，gorm 将选择所有字段
func (g *Gormx) Select(query interface{}, args ...interface{}) *gorm.DB {
	return g.client.Select(query, args...)
}

// Where 查询条件
func (g *Gormx) Where(query interface{}, args ...interface{}) *gorm.DB {
	return g.client.Where(query, args...)
}

// Raw 查询原生 SQL 构建器
func (g *Gormx) Raw(sql string, values ...interface{}) *gorm.DB {
	return g.client.Raw(sql, values...)
}

// Exec 执行原生 SQL 构建器
func (g *Gormx) Exec(sql string, values ...interface{}) *gorm.DB {
	return g.client.Exec(sql, values...)
}

var db *Gormx

// DB get gormx
func DB() *Gormx {
	return db
}

// Client get db client
func Client() *gorm.DB {
	return db.client
}

// Ping 判断服务是否可用
func Ping() error {
	return db.Ping()
}

// Table 定义表名
func Table(name string, args ...interface{}) *gorm.DB {
	return db.Table(name, args...)
}

// Model 指定操作模型
func Model(value interface{}) *gorm.DB {
	return db.Model(value)
}

// Create 创建数据
func Create(value interface{}) error {
	return db.Create(value)
}

// First 获取第一条记录（主键升序），且没有找到记录时，它会返回 gorm.ErrRecordNotFound 错误
func First(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.First(dest, conds...)
}

// Find 获取多条记录（查找所有符合给定条件的记录）
func Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return db.Find(dest, conds...)
}

// Select 选择指定查询、创建、更新时所需的字段，默认情况下，gorm 将选择所有字段
func Select(query interface{}, args ...interface{}) *gorm.DB {
	return db.Select(query, args...)
}

// Where 查询条件
func Where(query interface{}, args ...interface{}) *gorm.DB {
	return db.Where(query, args...)
}

// Raw 查询原生 SQL 构建器
func Raw(sql string, values ...interface{}) *gorm.DB {
	return db.Raw(sql, values...)
}

// Exec 执行原生 SQL 构建器
func Exec(sql string, values ...interface{}) *gorm.DB {
	return db.Exec(sql, values...)
}

// Connect 连接数据库
func Connect(c *Config) (*Gormx, error) {

	defaultLogger := logger.New(NewWriter(stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
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

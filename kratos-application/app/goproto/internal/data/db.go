package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"kratos-application/app/goproto/internal/conf"
)

var ProviderSet = wire.NewSet(NewDB)

type DB struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(c *conf.Data, logger log.Logger) (*DB, error) {
	return nil, nil
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, err
	}

	return &DB{
		db:  db,
		log: log.NewHelper(logger),
	}, nil
}

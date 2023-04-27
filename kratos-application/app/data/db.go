package data

import "github.com/go-kratos/kratos/v2/log"

type Data struct {
	log *log.Helper

	User *User
}

func NewData(logger log.Logger) *Data {
	return &Data{
		log: log.NewHelper(logger),

		User: &User{},
	}
}

package data

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// UserTable 表名
const UserTable = "user"

// User 用户
type User struct {
	ID        uint64         `json:"id" gorm:"id"`                 // 用户ID
	Account   string         `json:"account" gorm:"account"`       // 账号
	Email     string         `json:"email" gorm:"email"`           // 邮箱
	Password  string         `json:"password" gorm:"password"`     // 密码
	PhoneArea string         `json:"phone_area" gorm:"phone_area"` // 手机国家电话区号
	Phone     string         `json:"phone" gorm:"phone"`           // 手机号码
	Uname     string         `json:"uname" gorm:"uname"`           // 用户名
	Nickname  string         `json:"nickname" gorm:"nickname"`     // 昵称
	Sex       int8           `json:"sex" gorm:"sex"`               // 性别 1=男 2=女
	Birthday  sql.NullTime   `json:"birthday" gorm:"birthday"`     // 生日(年-月-日)
	AvatarUrl string         `json:"avatar_url" gorm:"avatar_url"` // 头像地址
	CoverUrl  string         `json:"cover_url" gorm:"cover_url"`   // 封面地址
	CreatedAt time.Time      `json:"created_at" gorm:"created_at"` // 创建时间
	UpdatedAt time.Time      `json:"updated_at" gorm:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"` // 删除时间
}

func (*User) TableName() string {
	return UserTable
}

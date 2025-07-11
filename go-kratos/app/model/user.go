package model

import (
	"time"
)

// UserTypeId 用户类型
type UserType = int

const (
	UserTypeUnknown  UserType = 0 //
	UserTypeOrdinary UserType = 1 // 普通用户
)

// UserTable 用户表名
const UserTable = "user"

// User 用户
type User struct {
	ID        int       `gorm:"column:id;type:int;autoIncrement;primaryKey"`
	TypeId    UserType  `gorm:"column:type_id;type:int;not null;default:0;comment:类型"`
	Account   string    `gorm:"column:account;type:string;size:120;not null;index;default:'';comment:账号"`
	Mobile    string    `gorm:"column:mobile;type:string;size:120;not null;index;default:'';comment:手机号"`
	Email     string    `gorm:"column:email;type:string;size:120;not null;index;default:'';comment:邮箱"`
	Password  string    `gorm:"column:password;type:string;size:120;not null;default:'';comment:密码"`
	Nickname  string    `gorm:"column:nickname;type:string;size:120;not null;default:'';comment:昵称"`
	AvatarUrl string    `gorm:"column:avatar_url;type:string;size:255;not null;default:'';comment:头像地址"`
	Sex       Sex       `gorm:"column:sex;type:int;not null;default:0;comment:性别 1男 2女"`
	States    State     `gorm:"column:states;type:int;not null;default:0;comment:状态 0封禁 1正常"`
	CreatedAt time.Time `gorm:"column:created_at;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;comment:更新时间"`
}

func (m *User) TableName() string {
	return UserTable
}

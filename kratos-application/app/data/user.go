package data

import (
	"time"

	"kratos-application/core/gormx"
)

const UserTable = "user"

// User 用户表
type User struct {
	ID        uint32    `gorm:"column:id" json:"id"`
	Uname     string    `gorm:"column:uname" json:"uname"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*User) TableName() string {
	return UserTable
}

func (*User) Create(user *User) error {
	return gormx.Create(user)
}

func (*User) FirstEmail(email string) (User, error) {
	var user User
	err := gormx.Where("email = ?", email).First(&user).Error
	return user, err
}

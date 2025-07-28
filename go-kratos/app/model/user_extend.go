package model

const UserExtendTable = "user_extend"

// UserExtend 用户扩展
type UserExtend struct {
	ID       int `gorm:"column:id;type:int;autoIncrement;primaryKey"`
	Uid      int `gorm:"column:uid;type:int;not null;index;comment:用户UID"`
	Birthday int `gorm:"column:birthday;type:int;not null;default:0;comment:状态 生日(年月日)"`
	Height   int `gorm:"column:height;type:int;not null;default:0;comment:身高(cm)"`
	Weight   int `gorm:"column:weight;type:int;not null;default:0;comment:体重(kg)"`
}

func (m *UserExtend) TableName() string {
	return UserExtendTable
}

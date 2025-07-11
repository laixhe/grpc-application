package model

const UserExtendTable = "user_extend"

// UserExtend 用户扩展
type UserExtend struct {
	ID            int    `gorm:"column:id;type:int;autoIncrement;primaryKey"`
	Uid           int    `gorm:"column:uid;type:int;not null;index;comment:用户UID"`
	WechatUnionid string `gorm:"column:wechat_unionid;type:string;size:255;not null;default:'';comment:微信unionid"`
	WechatOpenid  string `gorm:"column:wechat_openid;type:string;size:255;not null;default:'';index;comment:微信openid"`
}

func (m *UserExtend) TableName() string {
	return UserExtendTable
}

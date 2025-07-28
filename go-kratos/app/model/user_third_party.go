package model

const UserThirdPartyTable = "user_third_party"

// UserThirdParty 用户第三方
type UserThirdParty struct {
	ID            int    `gorm:"column:id;type:int;autoIncrement;primaryKey"`
	Uid           int    `gorm:"column:uid;type:int;not null;index;comment:用户UID"`
	WechatUnionid string `gorm:"column:wechat_unionid;type:string;size:200;not null;default:'';comment:微信unionid"`
	WechatOpenid  string `gorm:"column:wechat_openid;type:string;size:200;not null;index;default:'';comment:微信openid"`
}

func (m *UserThirdParty) TableName() string {
	return UserThirdPartyTable
}

package model

// UserSex 用户性别
type UserSex = int

const (
	UserSexUnknown UserSex = 0 // 未填写
	UserSexMale    UserSex = 1 // 男
	UserSexFemale  UserSex = 2 // 女
)

func IsUserSexValid(s UserSex) bool {
	switch s {
	case UserSexMale, UserSexFemale:
		return true
	}
	return false
}

func GetUserSexText(s UserSex) string {
	switch s {
	case UserSexMale:
		return "男"
	case UserSexFemale:
		return "女"
	default:
		return ""
	}
}

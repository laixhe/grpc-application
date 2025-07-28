package model

// UserTypeId 用户类型
type UserType = int

const (
	UserTypeOrdinary UserType = 1 // 普通用户
)

func IsUserTypeValid(t UserType) bool {
	switch t {
	case UserTypeOrdinary:
		return true
	}
	return false
}

func GetUserTypeText(t UserType) string {
	switch t {
	case UserTypeOrdinary:
		return "普通用户"
	default:
		return ""
	}
}

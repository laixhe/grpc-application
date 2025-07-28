package model

// UserState 状态
type UserState = int

const (
	UserStateBanned UserState = 0 // 禁用
	UserStateNormal UserState = 1 // 正常
)

func IsUserStateValid(s UserState) bool {
	switch s {
	case UserStateBanned, UserStateNormal:
		return true
	}
	return false
}

func GetUserStateText(s UserState) string {
	switch s {
	case UserStateBanned:
		return "禁用"
	case UserStateNormal:
		return "正常"
	default:
		return ""
	}
}

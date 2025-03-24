package constant

type UserStatus int8

const (
	USER_STATUS_INACTIVED UserStatus = 0 + iota
	USER_STATUS_ACTIVED 
)

func (s UserStatus) Pointer() *UserStatus{
	return &s
}

func (s *UserStatus) Value() UserStatus {
	if s == nil {
		return 0
	}
	return *s
}

type UserAction string

const (
	USER_ACTION_LOGIN UserAction = "LOGIN"
	USER_ACTION_USERNAME UserAction = "USERNAME"
	USER_ACTION_PASSWORD UserAction = "PASSWORD"
)


func (s UserAction) Pointer() *UserAction {
	return &s
}

func (s *UserAction) Value() UserAction {
	if s == nil {
		return ""
	}

	return *s
}

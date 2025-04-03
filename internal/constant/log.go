package constant

type LogStatus string

const (
	LOG_STATUS_INACTIVED LogStatus = "FAIL"
	LOG_STATUS_ACTIVED LogStatus 	= "SUCCESS"
)

func (s LogStatus) Pointer() *LogStatus{
	return &s
}

func (s *LogStatus) Value() LogStatus {
	if s == nil {
		return ""
	}
	return *s
}
package usecase

import (
	"miniBilling/internal/repository"

)



type VoiceReport interface{
}

type VoiceReportUsecase struct {
	report repository.VoiceReport
}

func NewVoiceReportUsecase(voiceReport repository.VoiceReport) VoiceReport {
	return &VoiceReportUsecase{report: voiceReport}
}

func (uc *VoiceReportUsecase) Cdr()   {
	
}
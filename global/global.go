package global

import (
	"miniBilling/internal/config"
	"miniBilling/internal/pkg/bot"
	"miniBilling/internal/pkg/logger"
	"miniBilling/internal/pkg/mysql"
)

var(
	Config 		*config.Config
	Logger 		*logger.LoggerZap
	Bot 		*bot.TeleBot
	Billing 	*mysql.BillingStruct
	VoiceReport *mysql.VoiceReportStruct
	DCN 		*mysql.DCNStruct
	// Mongo	*
)
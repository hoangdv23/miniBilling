package mysql

import "gorm.io/gorm"

type BillingStruct struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`

	Billing         string `mapstructure:"billing"`

	MaxIdleConns    int `mapstructure:"maxIdleConns"`
	MinOpenConns    int `mapstructure:"minOpenConns"`
	MaxOpenConns    int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`

	DB              *gorm.DB 
}


type DCNStruct struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`

	DCN             string `mapstructure:"DCN"`

	MaxIdleConns    int `mapstructure:"maxIdleConns"`
	MinOpenConns    int `mapstructure:"minOpenConns"`
	MaxOpenConns    int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`

	DB              *gorm.DB 
}

type VoiceReportStruct struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`

	VoiceReport     string `mapstructure:"VoiceReport"`

	MaxIdleConns    int `mapstructure:"maxIdleConns"`
	MinOpenConns    int `mapstructure:"minOpenConns"`
	MaxOpenConns    int `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int `mapstructure:"connMaxLifetime"`

	DB              *gorm.DB 
}
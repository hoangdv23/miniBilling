package config

type Config struct {
	AppName 		string 			`yaml:"app_name" mapstructure:"app_name"`
	HttpAddress 	int 			`yaml:"http_address" mapstructure:"http_address"`
	Bot 			Bot				`yaml:"bot" mapstructure:"bot"`
	Mysql			*Billing		`yaml:"mysql_billing" mapstructure:"mysql_billing"`
	Mysql136		*VoiceReport	`yaml:"mysql_136" mapstructure:"mysql_136"`
	Mongo_db		*Mongo			`yaml:"mongo" mapstructure:"mongo"`
	Log 			Logger			`yaml:"log" mapstructure:"log"`
}
type Bot struct {
	Token_bot 	string `yaml:"token_bot" mapstructure:"token_bot"`
}
type Billing struct {
	Host 		string 		`yaml:"host_billing" mapstructure:"host_billing"`
	Port 		int			`yaml:"port_billing" mapstructure:"port_billing"`
	Username 	string		`yaml:"username_billing" mapstructure:"username_billing"`
	Password 	string		`yaml:"password_billing" mapstructure:"password_billing"`

	Billing 	string		`yaml:"db_billing" mapstructure:"db_billing"`
	DCN 		string		`yaml:"db_DCN" mapstructure:"db_DCN"`

	MaxIdleConns 		int		`yaml:"maxIdleConns" mapstructure:"maxIdleConns"`
	MinOpenConns 		int		`yaml:"minOpenConns" mapstructure:"minOpenConns"`
	MaxOpenConns 		int		`yaml:"maxOpenConns" mapstructure:"maxOpenConns"`
	ConnMaxLifetime 	int		`yaml:"connMaxLifetime" mapstructure:"connMaxLifetime"`

}

type VoiceReport struct {
	Host 		string 		`yaml:"host_136" mapstructure:"host_136"`
	Port 		int			`yaml:"port_136" mapstructure:"port_136"`
	Username 	string		`yaml:"username_136" mapstructure:"username_136"`
	Password 	string		`yaml:"password_136" mapstructure:"password_136"`

	VoiceReport string		`yaml:"db_136" mapstructure:"db_136"`

	MaxIdleConns 		int		`yaml:"maxIdleConns" mapstructure:"maxIdleConns"`
	MinOpenConns 		int		`yaml:"minOpenConns" mapstructure:"minOpenConns"`
	MaxOpenConns 		int		`yaml:"maxOpenConns" mapstructure:"maxOpenConns"`
	ConnMaxLifetime 	int		`yaml:"connMaxLifetime" mapstructure:"connMaxLifetime"`
}
type Mongo struct {
	Url_mongo 	string  	`yaml:"url" mapstructure:"url"`
	DB_mongo 	string		`yaml:"db_name" mapstructure:"db_name"`
}

type Logger  struct {
	Log_level 	string		`yaml:"log_level" mapstructure:"log_level"`
	File_log	string		`yaml:"file_log_name" mapstructure:"file_log_name"`
	Max_size	int			`yaml:"max_size" mapstructure:"max_size"`
	Max_backup	int			`yaml:"max_backups" mapstructure:"max_backups"`
	Max_age		int			`yaml:"max_age" mapstructure:"max_age"`
	Compress	bool		`yaml:"compress" mapstructure:"compress"`
}
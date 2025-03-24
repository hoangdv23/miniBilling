package config
import(
	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error){
	cfg := &Config{}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}
	err = viper.Unmarshal(&cfg)
	if err !=  nil {
		return cfg, err
	}
	return cfg, err
}
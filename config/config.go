package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type ConfigType struct {
	Porduction             bool          `mapstrucure:"PRODUCTION"`
	DbHost                 string        `mapstructure:"HOST"`
	DbPort                 string        `mapstructure:"DB_PORT"`
	DbUser                 string        `mapstructure:"DB_USER"`
	DbPassword             string        `mapstructure:"DB_PASSWORD"`
	Port                   string        `mapstructure:"SERVER_PORT"`
	FrontEndOrigin         string        `mapstructure:"FRONTEND_ORIGIN"`
	JWTTokenSecret         string        `mapstructure:"JWT_SECRET"`
	TokenExpiresIn         time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge            time.Duration `mapstructure:"TOKEN_MAXAGE"`
	GoogleClientID         string        `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleClientSecret     string        `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOAuthRedirectUrl string        `mapstructure:"GOOGLE_OAUTH_REDIRECT_URL"`
}

// @USAGE
// config, _ := inits.LoadConfig(".")
func LoadEnvs(path string) (config ConfigType, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = viper.Unmarshal(&config)
	return
}

var Config, _ = LoadEnvs(".")

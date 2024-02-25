package config

import (
	"os"
	"time"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBUrl          string `mapstructure:"DATABASE_URL"`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	AwsRegion          string `mapstructure:"AWS_REGION"`
	AwsAccessKeyId     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`

	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     string `mapstructure:"SMTP_PORT"`
	SmtpUsername string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`

	Port string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

// func LoadConfig() (config Config, err error) {
// 	viper.SetConfigFile("app.env")
// 	viper.ReadInConfig()
// 	viper.AutomaticEnv()

// 	err = viper.Unmarshal(&config)
// 	return
// }

func LoadConfig() (config Config, err error) {
	return Config{
		DBUserName:         os.Getenv("POSTGRES_USER"),
		DBHost:             os.Getenv("POSTGRES_HOST"),
		DBUserPassword:     os.Getenv("POSTGRES_PASSWORD"),
		DBName:             os.Getenv("POSTGRES_DB"),
		DBPort:             os.Getenv("POSTGRES_PORT"),
		DBUrl:              os.Getenv("DATABASE_URL"),
		JwtSecret:          os.Getenv("JWT_SECRET"),
		Port:               os.Getenv("PORT"),
		AwsRegion:          os.Getenv("AWS_REGION"),
		AwsAccessKeyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SmtpHost:           os.Getenv("SMTP_HOST"),
		SmtpPort:           os.Getenv("SMTP_PORT"),
		SmtpUsername:       os.Getenv("SMTP_USERNAME"),
		SmtpPassword:       os.Getenv("SMTP_PASSWORD"),
	}, nil
}

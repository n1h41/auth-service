package config

import "github.com/spf13/viper"

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
  SmtpHost   string `mapstructure:"SMTP_HOST"`
  SmtpPort   string `mapstructure:"SMTP_PORT"`
  SmtpUser   string `mapstructure:"SMTP_USER"`
  SmtpPass   string `mapstructure:"SMTP_PASS"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	viper.Unmarshal(&config)
	return
}

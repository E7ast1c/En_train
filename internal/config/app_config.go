package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	DB  DBConfig
	Api ApiConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type ApiConfig struct {
	Port string
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func GetAppConfig() (*AppConfig, error) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		return nil, err
	}

	apiConfig := ApiConfig{Port: viper.GetString("api.port")}

	dbConfig := DBConfig{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		Username: os.Getenv("username"),
		DBName:   os.Getenv("dbname"),
		SSLMode:  os.Getenv("sslmode"),
		Password: os.Getenv("db_password"),
	}

	return &AppConfig{
		DB:  dbConfig,
		Api: apiConfig,
	}, nil
}

package config

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBUrl  string
	Host   string
	Port   string
	DBName string
	User   string
}

var AppConfig *Config

func LoadConfig() {
	_ = godotenv.Load()
	viper.AutomaticEnv()

	AppConfig = &Config{
		DBUrl: viper.GetString("DATABASE_URL"),
        Host:  viper.GetString("DB_HOST"),
        Port:  viper.GetString("DB_PORT"),
        DBName: viper.GetString("DB_NAME"),
        User:  viper.GetString("DB_USER"),
	}
	if AppConfig.DBUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}
    if AppConfig.Host == "" {
        log.Fatal("DB_HOST is not set")
    }
    if AppConfig.Port == "" {
        log.Fatal("DB_PORT is not set")
    }
    if AppConfig.DBName == "" {
        log.Fatal("DB_NAME is not set")
    }
    if AppConfig.User == "" {
        log.Fatal("DB_USER is not set")
    }
    log.Println("✅ Configuration loaded successfully")
	log.Println("Config loaded")
}
package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type AppConfig struct {
	DbConfig *DbConfig
}

func (config *AppConfig) LoadConfig() error {
	godotenv.Load(".env")
	dbconfig := &DbConfig{}
	dbconfig.Host = os.Getenv("DB_HOST")
	sPort := os.Getenv("DB_PORT")
	port_, err := strconv.Atoi(sPort)
	if err != nil {
		return err
	}

	dbconfig.Port = port_
	dbconfig.Database = os.Getenv("POSTGRES_DB")
	dbconfig.User = os.Getenv("POSTGRES_USER")
	dbconfig.Password = os.Getenv("POSTGRES_PASSWORD")
	config.DbConfig = dbconfig
	return nil
}

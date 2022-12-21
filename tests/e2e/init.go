package e2e

import (
	"os"
	"strconv"
	"universe-auth/config"
	"universe-auth/internal/auth"
	"universe-auth/internal/db"

	"github.com/joho/godotenv"
)

var worker *auth.AuthManager

func Initialize(con *db.Connection) *auth.AuthManager {
	worker, err := auth.NewAuthManager(con)
	if err != nil {
		panic(err)
	}
	return worker
}

func LoadEnv() (*config.DbConfig, error) {
	godotenv.Load(".env")
	dbconfig := &config.DbConfig{}
	dbconfig.Host = os.Getenv("DB_HOST")
	sPort := os.Getenv("DB_PORT")
	port_, err := strconv.Atoi(sPort)
	if err != nil {
		return nil, err
	}

	dbconfig.Port = port_
	dbconfig.Database = os.Getenv("POSTGRES_DB")
	dbconfig.User = os.Getenv("POSTGRES_USER")
	dbconfig.Password = os.Getenv("POSTGRES_PASSWORD")

	return dbconfig, nil
}

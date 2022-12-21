package main

import (
	"universe-auth/config"
	"universe-auth/internal/auth"
	"universe-auth/internal/db"
	"universe-auth/internal/pb"
)

func main() {
	config := config.AppConfig{}
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	con, err := db.New(*config.DbConfig)

	if err != nil {
		panic(err)
	}

	Initialize(con)
}

func Initialize(con *db.Connection) {
	worker, err := auth.NewAuthManager(con)

	if err != nil {
		panic(err)
	}

	pb.NewServer(worker)
}

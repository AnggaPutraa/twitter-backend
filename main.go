package main

import (
	"github.com/AnggaPutraa/twitter-backend/api"
	"github.com/AnggaPutraa/twitter-backend/configs"
	db "github.com/AnggaPutraa/twitter-backend/db/sqlc"
)

func main() {
	configuration := configs.LoadConfig()
	database := configs.OpenConnection(configuration.DBUrl)
	query := db.New(database)
	api.RunServer(configuration, query)
}

package main

import (
	"fmt"

	"github.com/AnggaPutraa/twitter-backend/api"
	"github.com/AnggaPutraa/twitter-backend/configs"
	db "github.com/AnggaPutraa/twitter-backend/db/sqlc"
)

func main() {
	configuration := configs.LoadConfig()
	fmt.Println(configuration.DBUrl)
	database := configs.OpenConnection(configuration.DBUrl)
	query := db.New(database)
	api.RunServer(configuration, query)
}

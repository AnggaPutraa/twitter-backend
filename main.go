package main

import (
	"fmt"
	"log"

	"github.com/AnggaPutraa/twitter-backend/configs"
)

func main() {
	configuration, err := configs.LoadConfig()

	if err != nil {
		log.Fatal("Can't read the env configuration")
	}

	fmt.Println(configuration.DBUrl)
}

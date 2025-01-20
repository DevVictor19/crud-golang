package main

import (
	"crud-golang/src/config"
	"crud-golang/src/database"
	user_model "crud-golang/src/models"
	"crud-golang/src/routers"
)

func main() {
	// config init
	config.LoadEnv()

	// database init
	database.InitMysqlConnection()
	database.AutoMigrate(&user_model.User{})

	// routers init
	routers.InitServerRoutes()
}

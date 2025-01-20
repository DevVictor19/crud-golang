package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

func GetConnection() *gorm.DB {
	return connection
}

type urlParams struct {
	user     string
	password string
	database string
	host     string
	port     string
}

func createConnectionUrl(params urlParams) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		params.user,
		params.password,
		params.host,
		params.port,
		params.database,
	)
}

func InitMysqlConnection() {
	dsn := createConnectionUrl(urlParams{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		database: os.Getenv("DB_DATABASE"),
	})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	connection = db
}

func AutoMigrate(models ...interface{}) {
	connection.AutoMigrate(models...)
}

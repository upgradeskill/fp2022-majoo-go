package configs

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDBMySql() (*sql.DB, error) {

	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DBCon, err := sql.Open(driver, dsn)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return DBCon, err
}

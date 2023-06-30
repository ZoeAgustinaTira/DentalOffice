package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var dbConnectionURL string

func LoadConfigFromFile() {

	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	dbConnectionURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_URL"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))

}

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConnectionURL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

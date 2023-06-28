package main

import (
	"database/sql"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	db, _ := sql.Open("mysql", "root:1234@tcp(localhost:3306)/dentaloffice")
	eng := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	router := routes.NewRouter(eng, db)
	router.MapRoutes()
	err = eng.Run()
	if err != nil {
		panic(err)
	}
}

package main

import (
	"database/sql"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/dentaloffice")

	if err != nil {
		log.Fatal(err)
	}
	eng := gin.Default()
	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}

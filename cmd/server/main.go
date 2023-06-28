package main

import (
	"database/sql"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/dentaloffice")
	eng := gin.Default()

	router := routes.NewRouter(eng, db)
	router.MapRoutes()
	if err := eng.Run(); err != nil {
		panic(err)
	}
}

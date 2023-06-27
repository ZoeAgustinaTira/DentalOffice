package main

import (
	"database/sql"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := sql.Open("mysql", "sarasa")
	eng := gin.Default()

	router := routes.NewRouter(eng, db)
	router.MapRoutes()
	if err := eng.Run(); err != nil {
		panic(err)
	}
}

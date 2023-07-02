package main

import (
	"github.com/ZoeAgustinatira/DentalOffice/cmd/config"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadConfigFromFile()
	db, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	eng := gin.Default()
	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}

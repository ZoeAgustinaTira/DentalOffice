package main

import (
	"github.com/ZoeAgustinatira/DentalOffice/cmd/config"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// @title Dental Office - Final exam
// @version 1.0
// @description This API Handle Products.
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support
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

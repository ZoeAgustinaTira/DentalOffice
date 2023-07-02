package main

import (
	"github.com/ZoeAgustinatira/DentalOffice/cmd/config"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/routes"
	"github.com/ZoeAgustinatira/DentalOffice/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title Dental Office - Final exam
// @version 1.0
// @description This API Handle Products.
// @termsOfService https://github.com/ZoeAgustinaTira/DentalOffice/blob/ec8e4d4c7aefae4d1e215ec421d4b1593773840f/README.md

// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support
func main() {
	config.LoadConfigFromFile()
	db, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	eng := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}

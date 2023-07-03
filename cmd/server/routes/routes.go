package routes

import (
	"database/sql"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/middleware"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/handler"
	"github.com/ZoeAgustinatira/DentalOffice/internal/dentist"
	"github.com/ZoeAgustinatira/DentalOffice/internal/patient"
	"github.com/ZoeAgustinatira/DentalOffice/internal/shift"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{
		eng: eng,
		db:  db,
	}
}

func (r *router) MapRoutes() {
	r.rg = r.eng.Group("/dentaloffice")
	r.buildDentistRoutes()
	r.buildPatientRoutes()
	r.buildShiftRoutes()
}

func (r *router) buildDentistRoutes() {
	repo := dentist.NewRepository(r.db)
	service := dentist.NewService(repo)
	handler := handler.NewDentist(service)

	bg := r.rg.Group("/dentists")
	auth := r.rg.Group("/dentists", middleware.TokenAuthMiddleware())
	{
		bg.GET("/:id", handler.GetByID())
		auth.POST("/", handler.Create())
		auth.PUT("/:id", handler.UpdateAll())
		auth.PATCH("/:id", handler.Update())
		auth.DELETE("/:id", handler.Delete())
	}
}

func (r *router) buildPatientRoutes() {
	repo := patient.NewRepository(r.db)
	service := patient.NewService(repo)
	handler := handler.NewPatient(service)

	bg := r.rg.Group("/patients")
	auth := r.rg.Group("/patients", middleware.TokenAuthMiddleware())
	{
		bg.GET("/:id", handler.GetByID())
		auth.POST("/", handler.Create())
		auth.PUT("/:id", handler.Update())
		auth.PATCH("/:id", handler.Update())
		auth.DELETE("/:id", handler.Delete())
	}
}

func (r *router) buildShiftRoutes() {
	repo := shift.NewRepository(r.db)
	service := shift.NewService(repo)
	handler := handler.NewShift(service)

	bg := r.rg.Group("/shifts")
	auth := r.rg.Group("/shifts", middleware.TokenAuthMiddleware())
	{
		bg.GET("/:id", handler.GetByID())
		bg.GET("/bydni/:dni", handler.GetByDNI())
		auth.POST("/", handler.Create())
		auth.PUT("/:id", handler.Update())
		auth.PATCH("/:id", handler.Update())
		auth.DELETE("/:id", handler.Delete())
	}
}

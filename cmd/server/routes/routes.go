package routes

import (
	"database/sql"
	"github.com/ZoeAgustinatira/DentalOffice/cmd/server/handler"
	"github.com/ZoeAgustinatira/DentalOffice/internal/dentist"

	_ "github.com/ZoeAgustinatira/DentalOffice/internal/patient"
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
}

func (r *router) buildDentistRoutes() {
	repo := dentist.NewRepository(r.db)
	service := dentist.NewService(repo)
	handler := handler.NewDentist(service)

	bg := r.rg.Group("/dentists")
	{
		bg.POST("/", handler.Create())
		bg.GET("/:id", handler.GetByID())
		//bg.PUT("/:id", handler.Update())   //Ver bien naming
		bg.PATCH("/:id", handler.Update()) //Ver bien naming
		bg.DELETE("/:id", handler.Delete())
	}

}

func (r *router) buildPatientRoutes() {

	/*repo := patient.NewRepository(r.db)
	service := patient.NewService(repo)
	handler := handler.NewPatient(service)

	bg := r.rg.Group("/patients")
	{
		bg.POST("/", handler.Create())
		bg.GET("/:id", handler.GetByID())
		bg.PUT("/:id", handler.Update())   //Ver bien naming
		bg.PATCH("/:id", handler.Update()) //Ver bien naming
		bg.DELETE("/:id", handler.Delete())
	}*/

}

func (r *router) buildShiftRoutes() {
	repo := shift.NewRepository(r.db)
	service := shift.NewService(repo)
	handler := handler.NewShift(service)

	bg := r.rg.Group("/shifts")
	{
		bg.POST("/", handler.Create())
		bg.GET("/:id", handler.GetByID())
		//bg.PUT("/:id", handler.Update())   //Ver bien naming
		bg.PATCH("/:id", handler.Update()) //Ver bien naming
		bg.DELETE("/:id", handler.Delete())
	}
}

package handler

import (
	"fmt"
	"github.com/ZoeAgustinatira/DentalOffice/internal/dentist"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Dentist struct {
	dentistService dentist.Service
}

func NewDentist(d dentist.Service) *Dentist {
	return &Dentist{
		dentistService: d,
	}
}

func (d *Dentist) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		dentist, err := d.dentistService.GetByID(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		var d domain.Dentist
		if dentist == d {
			c.JSON(http.StatusNotFound, err.Error()) //404
		}

		c.JSON(http.StatusOK, dentist)
	}
}

func (d *Dentist) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Dentist
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		if req.Name == "" || req.Surname == "" || req.Enrollment == "" {
			c.JSON(http.StatusUnprocessableEntity, "error: ¡incomplete fields!") //422
			return
		}

		exist := d.dentistService.Exists(c, req.Enrollment)
		if exist {
			c.JSON(http.StatusConflict, "error: the dentist already exist") //409
			return
		}

		newDentist, err := d.dentistService.Save(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error()) //500
			return
		}

		c.JSON(http.StatusCreated, newDentist)
	}
}

func (d *Dentist) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Dentist
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		req.ID = id

		dentistUpdate, err := d.dentistService.Update(c, req)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		var d domain.Dentist

		if dentistUpdate == d {
			c.JSON(http.StatusNotFound, err.Error()) //404
		}

		c.JSON(http.StatusOK, dentistUpdate)
	}
}

func (d *Dentist) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		err = d.dentistService.Delete(c, id)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		c.JSON(http.StatusNotFound, fmt.Sprintf("dentist %d deleted ", id))

	}
}

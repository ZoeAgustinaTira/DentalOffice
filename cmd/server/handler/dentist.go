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

// GetByID DentistByID godoc
// @Summary Get Dentist by ID
// @Tags Dentists
// @Description get Dentist by ID
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /dentists/{id} [get]
func (d *Dentist) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		dentist, err := d.dentistService.GetByID(id)
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

// Create CreateDentist godoc
// @Summary Create dentist
// @Tags Dentists
// @Description create dentist
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param domain.Dentist body domain.Dentist true "Dentist to create"
// @Success 201 {object} web.Response "dentist successfully created"
// @Failure 400 {object} web.Response "bad request"
// @Failure 409 {object} web.Response "error: the dentist already exist"
// @Failure 422 {object} web.Response "error: ¡incomplete fields!"
// @Failure 500 {object} web.Response "error while saving"
// @Router /dentists [post]
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

		exist := d.dentistService.Exists(req.Enrollment)
		if exist {
			c.JSON(http.StatusConflict, "error: the dentist already exist") //409
			return
		}

		newDentist, err := d.dentistService.Save(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error()) //500
			return
		}

		c.JSON(http.StatusCreated, newDentist)
	}
}

// Update UpdateDentist godoc
// @Summary Update dentist
// @Tags Dentists
// @Description update dentist
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param id path int true "id"
// @Param domain.Dentist body domain.Dentist true "Dentist to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /dentists/{id} [patch]
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

		dentistUpdate, err := d.dentistService.Update(req)
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

// UpdateAll UpdateAllDentist godoc
// @Summary Update all dentist by id
// @Tags Dentists
// @Description update all dentist by id
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param id path int true "id"
// @Param domain.Dentist body domain.Dentist true "Dentist to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /dentists/{id} [put]
func (d *Dentist) UpdateAll() gin.HandlerFunc {
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

		var fields []string
		if req.Name == "" {
			fields = append(fields, "name")
		}

		if req.Surname == "" {
			fields = append(fields, "surname")
		}

		if req.Enrollment == "" {
			fields = append(fields, "enrollment")
		}
		if len(fields) != 0 {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("field is missing: %v", fields)) //400
		}

		dentistUpdate, err := d.dentistService.UpdateAll(req)
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

// Delete DeleteDentist godoc
// @Summary Delete dentist by id
// @Tags Dentists
// @Description delete dentist by id
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param id path int true "id"
// @Success 204 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /dentists/{id} [delete]
func (d *Dentist) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		err = d.dentistService.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		c.JSON(http.StatusNoContent, fmt.Sprintf("dentist %d deleted", id))

	}
}

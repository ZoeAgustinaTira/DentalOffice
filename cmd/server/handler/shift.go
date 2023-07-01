package handler

import (
	"fmt"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
	"github.com/ZoeAgustinatira/DentalOffice/internal/shift"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Shift struct {
	shiftService shift.Service
}

func NewShift(s shift.Service) *Shift {
	return &Shift{
		shiftService: s,
	}
}

// GetByID ShiftByID godoc
// @Summary Get Shift by ID
// @Tags Shifts
// @Description get shift by ID
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /shifts/{id} [get]
func (s *Shift) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		shift, err := s.shiftService.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		var sh domain.Shift
		if shift == sh {
			c.JSON(http.StatusNotFound, err.Error()) //404
		}

		c.JSON(http.StatusOK, shift)
	}
}

// GetByDNI ShiftByDNI godoc
// @Summary Get Shift by DNI patient
// @Tags Shifts
// @Description get shift by DNI patient
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /shifts/bydni/{dni} [get]
func (s *Shift) GetByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dni := c.Param("dni")

		shift, err := s.shiftService.GetByDNI(dni)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		var sh domain.Shift
		if shift == sh {
			c.JSON(http.StatusNotFound, err.Error()) //404
		}

		c.JSON(http.StatusOK, shift)
	}
}

// Create CreateShift godoc
// @Summary Create Shift
// @Tags Shifts
// @Description create shift
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param domain.Shift body domain.Shift true "Shift to create"
// @Success 201 {object} web.Response "shift successfully created"
// @Failure 400 {object} web.Response "bad request"
// @Failure 409 {object} web.Response "error: the shift already exist"
// @Failure 422 {object} web.Response "error: ¡incomplete fields!"
// @Failure 500 {object} web.Response "error while saving"
// @Router /shifts [post]
func (s *Shift) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Shift
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		if req.Date == "" || req.Time == "" || req.DentistID == 0 || req.PatientID == 0 {
			c.JSON(http.StatusUnprocessableEntity, "error: ¡incomplete fields!") //422
			return
		}

		exist := s.shiftService.Exist(req)
		if exist != nil {
			c.JSON(http.StatusConflict, exist.Error()) //409
			return
		}

		newShift, err := s.shiftService.Save(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error()) //500
			return
		}

		c.JSON(http.StatusCreated, newShift)
	}
}

// Update UpdateShift godoc
// @Summary Update Shift
// @Tags Shifts
// @Description update shift
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param id path int true "id"
// @Param domain.Shift body domain.Shift true "Shift to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /shifts/{id} [patch]
func (s *Shift) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Shift
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

		shiftUpdate, err := s.shiftService.Update(req)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		var sh domain.Shift

		if shiftUpdate == sh {
			c.JSON(http.StatusNotFound, err.Error()) //404
		}

		c.JSON(http.StatusOK, shiftUpdate)
	}
}

// UpdateAll UpdateAllShift godoc
// @Summary Update all Shift by id
// @Tags Shifts
// @Description update all shift by id
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param id path int true "id"
// @Param domain.Shift body domain.Shift true "Shift to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /shifts/{id} [put]
func (s *Shift) UpdateAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Shift
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
		if req.Date == "" {
			fields = append(fields, "date")
		}

		if req.Time == "" {
			fields = append(fields, "time")
		}

		if req.DentistID == 0 {
			fields = append(fields, "dentist_id")
		}
		if req.PatientID == 0 {
			fields = append(fields, "patient_id")
		}
		if len(fields) != 0 {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("field is missing: %v", fields)) //400
		}

		shiftUpdate, err := s.shiftService.UpdateAll(req)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		var sh domain.Shift

		if shiftUpdate == sh {
			c.JSON(http.StatusNotFound, err.Error()) //404
		}

		c.JSON(http.StatusOK, shiftUpdate)
	}
}

// Delete DeleteShift godoc
// @Summary Delete Shift by id
// @Tags Shifts
// @Description delete shift by id
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param id path int true "id"
// @Success 204 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /shifts/{id} [delete]
func (s *Shift) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		err = s.shiftService.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		c.JSON(http.StatusNoContent, fmt.Sprintf("shift %d deleted ", id))
	}
}

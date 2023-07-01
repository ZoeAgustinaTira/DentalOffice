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

		c.JSON(http.StatusNotFound, fmt.Sprintf("shift %d deleted ", id))

	}
}

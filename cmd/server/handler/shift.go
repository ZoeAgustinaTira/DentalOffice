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

		shift, err := s.shiftService.GetByID(c, id)
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

		/*exist := d.dentistService.Exists(c, req.Enrollment)
		if exist {
			c.JSON(http.StatusConflict, "error: the dentist already exist") //409
			return
		}*/

		newShift, err := s.shiftService.Save(c, req)
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

		shiftUpdate, err := s.shiftService.Update(c, req)
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

		err = s.shiftService.Delete(c, id)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		c.JSON(http.StatusNotFound, fmt.Sprintf("shift %d deleted ", id))

	}
}

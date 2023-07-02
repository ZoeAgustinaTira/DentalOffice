package handler

import (
	"fmt"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
	"github.com/ZoeAgustinatira/DentalOffice/internal/patient"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Patient struct {
	patientService patient.Service
}

func NewPatient(p patient.Service) *Patient {
	return &Patient{
		patientService: p,
	}
}

// GetByID PatientByID godoc
// @Summary Get Patient by ID
// @Tags Patients
// @Description get patient by ID
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /patients/{id} [get]
func (p *Patient) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("GetByID")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		patient, err := p.patientService.GetByID(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		var d domain.Patient
		if patient == d {
			c.JSON(http.StatusNotFound, err.Error()) //404
		}

		c.JSON(http.StatusOK, patient)
	}
}

// Create CreatePatient godoc
// @Summary Create Patient
// @Tags Patients
// @Description create patient
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param domain.Patient body domain.Patient true "Patient to create"
// @Success 201 {object} web.Response "patient successfully created"
// @Failure 400 {object} web.Response "bad request"
// @Failure 409 {object} web.Response "error: the patient already exist"
// @Failure 422 {object} web.Response "error: ¡incomplete fields!"
// @Failure 500 {object} web.Response "error while saving"
// @Router /patients [post]
func (p *Patient) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Patient
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		if req.Name == "" || req.Surname == "" || req.DNI == "" || req.Address == "" || req.DischargeDate == "" {
			c.JSON(http.StatusUnprocessableEntity, "error: ¡incomplete fields!") //422
			return
		}

		exist := p.patientService.Exists(req.DNI)
		if exist {
			c.JSON(http.StatusConflict, "error: the patient already exist") //409
			return
		}

		newPatient, err := p.patientService.Save(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error()) //500
			return
		}

		c.JSON(http.StatusCreated, newPatient)
	}
}

// Update UpdatePatient godoc
// @Summary Update Patient
// @Tags Patients
// @Description update patient
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param id path int true "id"
// @Param domain.Patient body domain.Patient true "Patient to update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /patients/{id} [patch]
func (p *Patient) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Patient
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

		patientUpdate, err := p.patientService.Update(req)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		var d domain.Patient

		if patientUpdate == d {
			c.JSON(http.StatusNotFound, err.Error()) //404
		}

		c.JSON(http.StatusOK, patientUpdate)
	}
}

// Delete DeletePatient godoc
// @Summary Delete Patient by id
// @Tags Patients
// @Description delete patient by id
// @Accept  json
// @Produce  json
// @Param token header int true "token"
// @Param id path int true "id"
// @Success 204 {object} web.Response
// @Failure 400 {object} web.Response "bad request"
// @Failure 404 {object} web.Response "not found"
// @Router /patients/{id} [delete]
func (p *Patient) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error()) //400
			return
		}

		err = p.patientService.Delete(c, id)
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error()) //404
			return
		}

		c.JSON(http.StatusNotFound, fmt.Sprintf("patient %d deleted ", id))

	}
}

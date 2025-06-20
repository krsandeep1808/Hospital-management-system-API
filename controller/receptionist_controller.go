package controller

import (
	"clinic-portal/model"
	"clinic-portal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePatient(c *gin.Context) {
	var p model.Patient
	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	store.Patients = append(store.Patients, p)
	c.JSON(http.StatusOK, p)
}

func GetPatients(c *gin.Context) {
	c.JSON(http.StatusOK, store.Patients)
}

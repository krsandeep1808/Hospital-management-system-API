package controller

import (
	"clinic-portal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewPatients(c *gin.Context) {
	c.JSON(http.StatusOK, store.Patients)
}

package routes

import (
    "clinic-portal/controller"
    "clinic-portal/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.POST("/login", controller.Login)

    reception := r.Group("/reception")
    reception.Use(middleware.AuthMiddleware("receptionist"))
    {
        reception.POST("/patients", controller.CreatePatient)
        reception.GET("/patients", controller.GetPatients)
    }

    doctor := r.Group("/doctor")
    doctor.Use(middleware.AuthMiddleware("doctor"))
    {
        doctor.GET("/patients", controller.ViewPatients)
    }
}
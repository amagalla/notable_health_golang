package routes

import (
	"notable_health/cmd/main/middleware"
	"notable_health/cmd/main/routes/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/appointments/allPhysicians", api.AllPhysicians)

	router.POST("/api/appointments/insertPhysicians", api.InsertPhysician)
	router.POST("/api/appointments/addAppointment/:IdPhysician", middleware.CheckTime, api.AddAppointment)

	router.DELETE("/api/appointments/cancelAppointment", api.CancelAppointment)
}

package routes

import (
	"notable_health/cmd/main/routes/api"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/appointments/allPhysicians", api.AllPhysicians)

	router.POST("/api/appointments/insertPhysicians", api.InsertPhysician)
}

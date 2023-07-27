package routes

import (
	"notable_health/cmd/main/routes/schedules"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/api/getPhysicians", schedules.GetPhysicians)

	router.POST("/api/insertPhysician", schedules.PostPhysician)
}

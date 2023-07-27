package routes

import (
	"notable_health/cmd/main/routes/schedules"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/insertPhysician", schedules.PostPhysicians)

	router.GET("/getPhysicianList", schedules.GetPhysicianList)
}

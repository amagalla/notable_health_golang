package schedules

import (
	"notable_health/cmd/main/controllers"
	"notable_health/cmd/main/models"

	"github.com/gin-gonic/gin"
)

func PostPhysician(c *gin.Context) {
	var physicians models.PostPhysiciansData

	if err := c.BindJSON(&physicians); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to parse request body",
		})

		return
	}

	if err := controllers.InsertPhysician(&physicians); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "physicians inserted into table",
	})
}

func GetPhysicians(c *gin.Context) {
	physicanListResp, err := controllers.GetPhysicianList()

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	resp := models.GetPhysicianResponse{
		PhysicianList: physicanListResp,
	}

	c.JSON(200, resp)
}

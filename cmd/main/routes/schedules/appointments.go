package schedules

import (
	"notable_health/cmd/main/controllers"
	"notable_health/cmd/main/models"

	"github.com/gin-gonic/gin"
)

func PostPhysicians(c *gin.Context) {
	// create a struct with expected values
	var physicians models.Physicians

	// parse data to struct
	if err := c.BindJSON(&physicians); err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to parse request body",
		})

		return
	}

	// pass struct as pointer to contoller
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

func GetPhysicianList(c *gin.Context) {
	getListResponse, err := controllers.GetPhysicians()

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	resp := models.PhysicianResponse{
		PhysicianList: getListResponse,
	}

	c.JSON(200, resp)
}

package api

import (
	"notable_health/cmd/main/controller"
	"notable_health/cmd/main/models"

	"github.com/gin-gonic/gin"
)

func AllPhysicians(c *gin.Context) {
	physicianListData, err := controller.GetPhysicianList()

	if err != nil {
		c.JSON(400, gin.H{
			"error": "erorr retrieving physician data",
		})

		return
	}

	resp := models.PhysicianListResponse{
		PhysicianList: physicianListData,
	}

	c.JSON(200, resp)
}

func InsertPhysician(c *gin.Context) {
	var reqBody models.InsertPhysicianData

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{
			"error": "error binding request data",
		})

		return
	}

	if err := controller.InsertPhysicianData(&reqBody); err != nil {
		c.JSON(400, gin.H{
			"error": "error inserting request data",
		})

		return
	}

	c.JSON(200, gin.H{
		"success": "physician inserted",
	})
}

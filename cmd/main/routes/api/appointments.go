package api

import (
	"notable_health/cmd/main/controller"
	"notable_health/cmd/main/models"
	"strconv"

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

func AppointmentLookup(c *gin.Context) {
	IdPhy, _ := strconv.Atoi(c.Query("IdPhysician"))
	date := c.Query("date_column")

	appointmentListData, err := controller.GetAppointments(IdPhy, date)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	resp := models.AppointmentListResponse{
		Appointments: appointmentListData,
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

	if err := controller.CheckValidPhysician(&reqBody); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
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

func AddAppointment(c *gin.Context) {
	idPhysician, _ := strconv.Atoi(c.Param("IdPhysician"))

	appointmentData, exists := c.Get("reqBody")
	if !exists {
		c.JSON(400, gin.H{
			"error": "reqBody not found in context",
		})
		return
	}

	reqBody := appointmentData.(models.AddAppointmentData)

	if err := controller.InsertAppointmentData(&reqBody, idPhysician); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"sucess": "appointment scheduled successfully",
	})
}

func CancelAppointment(c *gin.Context) {

	IdAppointment, _ := strconv.Atoi(c.Query("IdAppointment"))
	IdPhysician, _ := strconv.Atoi(c.Query("IdPhysician"))

	err := controller.DeleteAppointment(IdAppointment, IdPhysician)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"success": "appointment cancelled successfully",
	})
}

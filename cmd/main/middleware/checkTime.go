package middleware

import (
	"notable_health/cmd/main/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckTime(c *gin.Context) {

	var reqBody models.AddAppointmentData

	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{
			"error": "error binding json body in CheckTime middleware",
		})

		c.Abort()
		return
	}

	time := reqBody.Time_Column

	timeSlice := strings.Split(time, ":")

	seconds, _ := strconv.Atoi(timeSlice[1][:2])

	if seconds%15 != 0 {
		c.JSON(400, gin.H{
			"error": "invalid time",
		})

		c.Abort()
		return
	}

	c.Set("reqBody", reqBody)

	c.Next()
}

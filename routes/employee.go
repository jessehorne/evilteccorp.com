package routes

import (
	"evilteccorp.com/database"
	"evilteccorp.com/database/models"
	"evilteccorp.com/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEmployee(c *gin.Context) {
	// get authed user
	user, err := helper.GetUserFromRequest(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "something went wrong grabbing your user",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "nice try bud",
		})
		return
	}

	// get projects created by user
	var solutions []models.Solution
	result := database.GDB.Where("user_id = ?", user.ID).Find(&solutions)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "something went wrong",
		})
		return
	}

	var formattedProjects []string
	for _, s := range solutions {
		var p models.Project
		result := database.GDB.Find(&p, s.ProjectID)
		if result.RowsAffected != 0 {
			formattedProjects = append(formattedProjects, p.Title)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"evilcoin":  user.Coins,
		"solutions": result.RowsAffected,
		"projects":  formattedProjects,
	})
}

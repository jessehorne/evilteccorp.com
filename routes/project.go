package routes

import (
	"evilteccorp.com/database"
	"evilteccorp.com/database/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
)

type PostProjectRequest struct {
	Key string `json:"key" validate:"required"`

	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Answer      string `json:"answer" validate:"required"`
	Tags        string `json:"tags" validate:"required"`
	Reward      int    `json:"reward" validate:"required"`
}

func PostProject(c *gin.Context) {
	var body PostProjectRequest

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// validate request
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request",
			"errors": errors,
		})
		return
	}

	if body.Key != os.Getenv("PROJECT_KEY") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "nice try bud",
		})
		return
	}

	// create project
	result := database.GDB.Create(&models.Project{
		Title:       body.Title,
		Description: body.Description,
		Answer:      body.Answer,
		Tags:        body.Tags,
		Reward:      body.Reward,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "db err",
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

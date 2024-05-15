package routes

import (
	"evilteccorp.com/database"
	"evilteccorp.com/database/models"
	"evilteccorp.com/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type PostRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func PostRegister(c *gin.Context) {
	var body PostRegisterRequest

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

	hashed, err := helper.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong with bcrypt",
		})
		return
	}

	result := database.GDB.Create(&models.User{
		Email:    body.Email,
		Password: hashed,
		Coins:    0,
	})

	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong with gorm",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "You've registered an account. Good job! Visit '/api' for next steps.",
	})
}

type PostTokenRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func PostToken(c *gin.Context) {
	var body PostTokenRequest

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

	// get user by email to make sure it exists
	var user models.User
	tx := database.GDB.First(&user, "email = ?", body.Email)
	if tx.RowsAffected == 0 || tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credentials or error",
		})
		return
	}

	// check password
	if !helper.CheckPassword(user.Password, body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// generate new token
	tok := helper.MakeToken()
	updated := time.Now()

	user.Token = &tok
	user.TokenUpdatedAt = &updated
	database.GDB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"token": tok,
	})
}

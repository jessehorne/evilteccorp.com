package helper

import (
	"crypto/rand"
	"errors"
	"evilteccorp.com/database"
	"evilteccorp.com/database/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func MakeToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", b)
}

func GetUserFromRequest(c *gin.Context) (*models.User, error) {
	token := c.GetHeader("Authorization")

	if token == "" {
		return nil, errors.New("nil token")
	}

	var user *models.User
	result := database.GDB.
		Where("token = ?", token).
		Where("token_updated_at > date_add(NOW(), interval -1 hour)").
		First(&user)

	if result.RowsAffected == 0 || result.Error != nil {
		return nil, errors.New("query error")
	}

	return user, nil
}

package response

import (
	"belajar-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserResponseSingle struct {
	Status  int
	Message string
	Data    UserResponseData
}

type UserResponseData struct {
	ID          int
	Image       string
	Name        string
	Email       string
	Phone       string
	DateOfBirth time.Time
	CityID      int
	ProvinceID  int
	Address     string
	Password    string
	Token       string
	LastLogin   time.Time
}

func UserResponse1(c *gin.Context, msg string, token string, user models.User) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.JSON(http.StatusOK, UserResponseSingle{
		Status:  http.StatusOK,
		Message: msg,
		Data: UserResponseData{
			ID:          user.ID,
			Image:       user.Image,
			Name:        user.Name,
			Email:       user.Email,
			Phone:       user.Phone,
			DateOfBirth: user.DateOfBirth,
			CityID:      user.CityID,
			ProvinceID:  user.ProvinceID,
			Address:     user.Address,
			Password:    user.Password,
			Token:       token,
			LastLogin:   user.LastLogin,
		},
	})

}

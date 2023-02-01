package response

import (
	"belajar-api/helper"
	"belajar-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RatingResponse struct {
	Status    int
	Message   string
	TotalData int
	OrderName string
	OrderBy   string
	Page      int
	Limit     int
	Data      []RatingData
}

type RatingResponseSingle struct {
	Status  int
	Message string
	Data    RatingData
}

type RatingData struct {
	ID        int
	BookID    int
	UserID    int
	Rating    int
	Review    string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      UserResponseData
}

func RatingResponseList(c *gin.Context, msg string, ratings []models.Rating, orderName string, orderBy string, pagination helper.Pagination) {
	var ratingsResponse []RatingData
	for _, r := range ratings {
		ratingsResponse = append(ratingsResponse, RatingData{
			ID:        r.ID,
			BookID:    r.BookID,
			UserID:    r.UserID,
			Rating:    r.Rating,
			Review:    r.Review,
			CreatedAt: r.CreatedAt,
			User: UserResponseData{
				Name:  r.User.Name,
				Email: r.User.Email,
			},
		})
	}

	c.JSON(http.StatusOK, RatingResponse{
		Status:    http.StatusOK,
		Message:   msg,
		TotalData: len(ratingsResponse),
		OrderName: orderName,
		OrderBy:   orderBy,
		Page:      pagination.Page,
		Limit:     pagination.Limit,
		Data:      ratingsResponse,
	})
}

func RatingResponse1(c *gin.Context, msg string, r models.Rating) {
	c.JSON(http.StatusOK, RatingResponseSingle{
		Status:  http.StatusOK,
		Message: msg,
		Data: RatingData{
			ID:        r.ID,
			BookID:    r.BookID,
			UserID:    r.UserID,
			Rating:    r.Rating,
			Review:    r.Review,
			CreatedAt: r.CreatedAt,
			User: UserResponseData{
				ID:    r.User.ID,
				Name:  r.User.Name,
				Email: r.User.Email,
			},
		},
	})
}

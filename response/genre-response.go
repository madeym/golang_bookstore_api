package response

import (
	"belajar-api/helper"
	"belajar-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreResponse struct {
	Status    int
	Message   string
	TotalData int
	OrderName string
	OrderBy   string
	Page      int
	Limit     int
	Data      []GenreResponseData
}

type GenreResponseSingle struct {
	Status  int
	Message string
	Data    GenreResponseData
}

type GenreResponseData struct {
	ID          int
	Name        string
	Description string
}

func GenreResponseList(c *gin.Context, msg string, genres []models.Genre, orderName string, orderBy string, pagination helper.Pagination) {
	var genresResponse []GenreResponseData

	for _, b := range genres {
		genresResponse = append(genresResponse, GenreResponseData{
			ID:          b.ID,
			Name:        b.Name,
			Description: b.Description,
		})
	}

	c.JSON(http.StatusOK,
		GenreResponse{
			Status:    http.StatusOK,
			Message:   "Genre Data Successfully Displayed",
			TotalData: len(genres),
			OrderName: orderName,
			OrderBy:   orderBy,
			Page:      pagination.Page,
			Limit:     pagination.Limit,
			Data:      genresResponse,
		},
	)
}

func GenreResponse1(c *gin.Context, msg string, genre models.Genre) {
	c.JSON(http.StatusOK, GenreResponseSingle{
		Status:  http.StatusOK,
		Message: msg,
		Data: GenreResponseData{
			ID:          genre.ID,
			Name:        genre.Name,
			Description: genre.Description,
		},
	})
}

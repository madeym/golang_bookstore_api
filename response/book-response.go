package response

import (
	"belajar-api/helper"
	"belajar-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BookResponse struct {
	Status    int
	Message   string
	TotalData int
	OrderName string
	OrderBy   string
	Page      int
	Limit     int
	Data      []ResponseData
}

type BookResponseSingle struct {
	Status  int
	Message string
	Data    ResponseData
}

type ResponseData struct {
	ID          int
	CoverImage  string
	Title       string
	Author      string
	PublishDate time.Time
	Page        int
	Description string
	Price       int
	Rating      int
	Discount    int
	Quantity    int
	GenreID     int
	StatusID    int
	Genre       GenreResponseData
	Status      StatusResponseData
}

func BookResponseList(c *gin.Context, msg string, books []models.Book, orderName string, orderBy string, pagination helper.Pagination) {
	var booksResponse []ResponseData
	for _, b := range books {
		booksResponse = append(booksResponse, ResponseData{
			ID:          b.ID,
			CoverImage:  b.CoverImage,
			Title:       b.Title,
			Author:      b.Author,
			PublishDate: b.PublishDate,
			Page:        b.Page,
			Price:       b.Price,
			Description: b.Description,
			Rating:      b.Rating,
			Discount:    b.Discount,
			Quantity:    b.Quantity,
			GenreID:     b.GenreID,
			StatusID:    b.StatusID,
			Genre: GenreResponseData{
				ID:          b.Genre.ID,
				Name:        b.Genre.Name,
				Description: b.Genre.Description,
			},
			Status: StatusResponseData{
				ID:   b.Status.ID,
				Name: b.Status.Name,
			},
		})
	}

	c.JSON(http.StatusOK,
		BookResponse{
			Status:    http.StatusOK,
			Message:   msg,
			TotalData: len(booksResponse),
			OrderName: orderName,
			OrderBy:   orderBy,
			Page:      pagination.Page,
			Limit:     pagination.Limit,
			Data:      booksResponse,
		},
	)
}

func BookResponse1(c *gin.Context, msg string, b models.Book) {
	c.JSON(http.StatusOK, BookResponseSingle{
		Status:  http.StatusOK,
		Message: msg,
		Data: ResponseData{
			ID:          b.ID,
			CoverImage:  b.CoverImage,
			Title:       b.Title,
			Author:      b.Author,
			PublishDate: b.PublishDate,
			Page:        b.Page,
			Price:       b.Price,
			Description: b.Description,
			Rating:      b.Rating,
			Discount:    b.Discount,
			Quantity:    b.Quantity,
			GenreID:     b.GenreID,
			StatusID:    b.StatusID,
			Genre: GenreResponseData{
				ID:          b.Genre.ID,
				Name:        b.Genre.Name,
				Description: b.Genre.Description,
			},
			Status: StatusResponseData{
				ID:   b.Status.ID,
				Name: b.Status.Name,
			},
		},
	})
}

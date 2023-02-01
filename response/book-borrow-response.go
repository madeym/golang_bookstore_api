package response

import (
	"belajar-api/helper"
	"belajar-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BookBorrowResponse struct {
	Status    int
	Message   string
	TotalData int
	Page      int
	Limit     int
	Data      []BookBorrowResponseData
}

type BookBorrowResponseSingle struct {
	Status  int
	Message string
	Data    BookBorrowResponseData
}

type BookBorrowResponseData struct {
	ID               int
	UserID           int
	User             UserResponseData
	BookID           int
	Book             ResponseData
	BorrowTime       time.Time
	ExpireBorrowTime time.Time
}

func BookBorrowResponseList(c *gin.Context, msg string, bb []models.BookBorrowScan, pagination helper.Pagination) {
	var bbResponse []BookBorrowResponseData
	for _, b := range bb {
		bbResponse = append(bbResponse, BookBorrowResponseData{
			ID:     b.ID,
			UserID: b.UserID,
			User: UserResponseData{
				ID:    b.User.ID,
				Name:  b.User.Name,
				Email: b.User.Email,
			},
			BookID: b.BookID,
			Book: ResponseData{
				ID:          b.BookID,
				Title:       b.Book.Title,
				Price:       b.Book.Price,
				Description: b.Book.Description,
				Rating:      b.Book.Rating,
				Discount:    b.Book.Discount,
				Quantity:    b.Book.Quantity,
				GenreID:     b.Book.GenreID,
				StatusID:    b.Book.StatusID,
				Genre: GenreResponseData{
					ID:          b.Book.Genre.ID,
					Name:        b.Book.Genre.Name,
					Description: b.Book.Genre.Description,
				},
				Status: StatusResponseData{
					ID:   b.Book.Status.ID,
					Name: b.Book.Status.Name,
				},
			},
			BorrowTime:       b.BorrowTime,
			ExpireBorrowTime: b.ExpireBorrowTime,
		})
	}

	c.JSON(http.StatusOK, BookBorrowResponse{
		Status:    http.StatusOK,
		Message:   msg,
		TotalData: len(bb),
		Page:      pagination.Page,
		Limit:     pagination.Limit,
		Data:      bbResponse,
	})
}

func BookBorrowResponse1(c *gin.Context, msg string, b models.BookBorrowScan) {
	c.JSON(http.StatusOK, BookBorrowResponseSingle{
		Status:  http.StatusOK,
		Message: msg,
		Data: BookBorrowResponseData{
			ID:     b.ID,
			UserID: b.UserID,
			User: UserResponseData{
				ID:    b.User.ID,
				Name:  b.User.Name,
				Email: b.User.Email,
			},
			BookID: b.BookID,
			Book: ResponseData{
				ID:          b.BookID,
				Title:       b.Book.Title,
				Price:       b.Book.Price,
				Description: b.Book.Description,
				Rating:      b.Book.Rating,
				Discount:    b.Book.Discount,
				Quantity:    b.Book.Quantity,
				GenreID:     b.Book.GenreID,
				StatusID:    b.Book.StatusID,
				Genre: GenreResponseData{
					ID:          b.Book.Genre.ID,
					Name:        b.Book.Genre.Name,
					Description: b.Book.Genre.Description,
				},
				Status: StatusResponseData{
					ID:   b.Book.Status.ID,
					Name: b.Book.Status.Name,
				},
			},
			BorrowTime:       b.BorrowTime,
			ExpireBorrowTime: b.ExpireBorrowTime,
		},
	})
}

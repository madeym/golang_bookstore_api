package response

import (
	"belajar-api/helper"
	"belajar-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionResponse struct {
	Status    int
	Message   string
	TotalData int
	Page      int
	Limit     int
	Data      []TransactionResponseData
}

type TransactionResponseSingle struct {
	Status  int
	Message string
	Data    TransactionResponseData
}

type TransactionResponseData struct {
	ID                 int
	UserID             int
	User               UserResponseData
	BookID             int
	Book               ResponseData
	Quantity           int
	PaymentBankID      int
	PaymentBank        BankResponseData
	PaymentBankAccount string
	PaymentDate        time.Time
}

func TransactionResponseList(c *gin.Context, msg string, transactions []models.TransactionScan, pagination helper.Pagination) {
	var tResponse []TransactionResponseData

	for _, t := range transactions {
		tResponse = append(tResponse, TransactionResponseData{
			ID:     t.ID,
			UserID: t.UserID,
			User: UserResponseData{
				ID:    t.User.ID,
				Name:  t.User.Name,
				Email: t.User.Email,
			},
			BookID: t.BookID,
			Book: ResponseData{
				ID:          t.Book.ID,
				Title:       t.Book.Title,
				Price:       t.Book.Price,
				Description: t.Book.Description,
				Rating:      t.Book.Rating,
				Discount:    t.Book.Discount,
				Quantity:    t.Book.Quantity,
				GenreID:     t.Book.GenreID,
				StatusID:    t.Book.StatusID,
				Genre: GenreResponseData{
					ID:          t.Book.Genre.ID,
					Name:        t.Book.Genre.Name,
					Description: t.Book.Genre.Description,
				},
				Status: StatusResponseData{
					ID:   t.Book.Status.ID,
					Name: t.Book.Status.Name,
				},
			},
			Quantity:      t.Quantity,
			PaymentBankID: t.PaymentBankID,
			PaymentBank: BankResponseData{
				ID:   t.PaymentBank.ID,
				Name: t.PaymentBank.Name,
			},
			PaymentBankAccount: t.PaymentBankAccount,
			PaymentDate:        t.PaymentDate,
		})
	}

	c.JSON(http.StatusOK, TransactionResponse{
		Status:    http.StatusOK,
		Message:   msg,
		TotalData: len(tResponse),
		Page:      pagination.Page,
		Limit:     pagination.Limit,
		Data:      tResponse,
	})
}

func TransactionResponse1(c *gin.Context, msg string, t *models.TransactionScan) {
	c.JSON(http.StatusOK, TransactionResponseSingle{
		Status:  http.StatusOK,
		Message: msg,
		Data: TransactionResponseData{
			ID:     t.ID,
			UserID: t.UserID,
			User: UserResponseData{
				ID:    t.User.ID,
				Name:  t.User.Name,
				Email: t.User.Email,
			},
			BookID: t.BookID,
			Book: ResponseData{
				ID:          t.Book.ID,
				Title:       t.Book.Title,
				Price:       t.Book.Price,
				Description: t.Book.Description,
				Rating:      t.Book.Rating,
				Discount:    t.Book.Discount,
				Quantity:    t.Book.Quantity,
				GenreID:     t.Book.GenreID,
				StatusID:    t.Book.StatusID,
				Genre: GenreResponseData{
					ID:          t.Book.Genre.ID,
					Name:        t.Book.Genre.Name,
					Description: t.Book.Genre.Description,
				},
				Status: StatusResponseData{
					ID:   t.Book.Status.ID,
					Name: t.Book.Status.Name,
				},
			},
			Quantity:      t.Quantity,
			PaymentBankID: t.PaymentBankID,
			PaymentBank: BankResponseData{
				ID:   t.PaymentBank.ID,
				Name: t.PaymentBank.Name,
			},
			PaymentBankAccount: t.PaymentBankAccount,
			PaymentDate:        t.PaymentDate,
		},
	})
}

package controllers

import (
	"belajar-api/helper"
	"belajar-api/models"
	"belajar-api/request"
	"belajar-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	keywordName, keyword := helper.GetKeywordQueryFromRequest(c)
	param := helper.GetParamQueryFromRequest(c, []string{"Title", "Price", "StatusID"}, []string{"title LIKE ?", "price < ?", "status_id = ?"}, []string{"", "", ""})
	orderName, orderBy := helper.GetOrderQueryFromRequest(c)
	pagination := helper.GeneratePagination(c)

	books, err := models.GetAllBooks(keywordName, keyword, param, orderName, orderBy, pagination)

	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.BookResponseList(c, "Book Data Successfully Displayed", books, orderName, orderBy, pagination)
}

func GetBookByID(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))

	b, err := models.FindBookByID(id)

	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.BookResponse1(c, "Book Data Successfully Displayed", b)
}

func CreateBook(c *gin.Context) {
	var BookRequest request.BookRequest

	err := c.ShouldBindJSON(&BookRequest)
	if helper.ShouldBindJSONError(c, err) {
		return
	}

	price, _ := BookRequest.Price.Int64()
	quantity, _ := BookRequest.Quantity.Int64()
	genreid, _ := BookRequest.GenreID.Int64()
	statusid, _ := BookRequest.StatusID.Int64()

	newbook := models.Book{
		Title:       BookRequest.Title,
		Price:       int(price),
		Description: BookRequest.Description,
		Rating:      BookRequest.Rating,
		Discount:    BookRequest.Discount,
		Quantity:    int(quantity),
		GenreID:     int(genreid),
		StatusID:    int(statusid),
	}

	b, err := models.CreateBook(newbook)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.BookResponse1(c, "Successfully Created New Book", b)
}

func UpdateBook(c *gin.Context) {
	var bookRequest request.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if helper.ShouldBindJSONError(c, err) {
		return
	}

	price, _ := bookRequest.Price.Int64()
	quantity, _ := bookRequest.Quantity.Int64()
	genreid, _ := bookRequest.GenreID.Int64()
	statusid, _ := bookRequest.StatusID.Int64()

	updateBook := models.Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
		Quantity:    int(quantity),
		GenreID:     int(genreid),
		StatusID:    int(statusid),
	}

	id, _ := helper.StringToInt(c.Param("id"))

	b, err := models.UpdateBook(id, updateBook)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.BookResponse1(c, "Successfully Updated Existing Book", b)
}

func DeleteBook(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))

	err := models.DeleteBook(id)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Successfully Deleted Existing Book",
	})
}

func BuyBook(c *gin.Context) {
	var request request.BuyBookRequest

	err := c.ShouldBindJSON(&request)
	if helper.ShouldBindJSONError(c, err) {
		return
	}

	id, _ := helper.StringToInt(c.Param("id"))

	t, _, msg, err := models.BuyBook(id, &request)

	if msg != "" || err != nil {
		helper.ResponseError(c, msg, err)
		return
	}

	response.TransactionResponse1(c, "Successfully Buy Book", &t)
}

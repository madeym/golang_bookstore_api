package controllers

import (
	"belajar-api/helper"
	"belajar-api/models"
	"belajar-api/response"
	"time"

	"github.com/gin-gonic/gin"
)

func GetBookBorrows(c *gin.Context) {
	keywordName, keyword := helper.GetKeywordQueryFromRequest(c)
	param := helper.GetParamQueryFromRequest(c, []string{"UserID", "IsExpired"}, []string{"user_id = ?", "expire_borrow_time < ?"}, []string{"", time.Now().String()})
	orderName, orderBy := helper.GetOrderQueryFromRequest(c)
	pagination := helper.GeneratePagination(c)

	b, err := models.GetBookBorrows(keywordName, keyword, param, orderName, orderBy, pagination)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.BookBorrowResponseList(c, "Book Borrow Data Successfully Displayed", b, pagination)
}

func BorrowBook(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))

	b, msg, err := models.CreateBorrowBook(id)
	if msg != "" || err != nil {
		helper.ResponseError(c, msg, err)
		return
	}

	response.BookBorrowResponse1(c, "Successfully Borrow Book", b)
}

func ReturnBookBorrow(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))

	b, msg, err := models.ReturnBookBorrow(id)
	if msg != "" || err != nil {
		helper.ResponseError(c, msg, err)
		return
	}

	response.BookBorrowResponse1(c, "Successfully Return Book", b)
}

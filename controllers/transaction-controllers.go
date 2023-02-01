package controllers

import (
	"belajar-api/helper"
	"belajar-api/models"
	"belajar-api/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	keywordName, keyword := helper.GetKeywordQueryFromRequest(c)
	param := helper.GetParamQueryFromRequest(c, []string{"UserID"}, []string{"user_id = ?"}, []string{""})
	fmt.Printf("param : %v : \n\n", param)
	pagination := helper.GeneratePagination(c)

	t, err := models.GetAllTransactions(pagination, param, keywordName, keyword)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.TransactionResponseList(c, "Transaction Data Successfully Displayed", t, pagination)
}

func GetTransactionByID(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))

	t, err := models.GetTransactionByID(id)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.TransactionResponse1(c, "Transaction Data Successfully Displayed", &t)
}

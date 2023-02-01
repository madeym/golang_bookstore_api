package controllers

import (
	"belajar-api/helper"
	"belajar-api/models"
	"belajar-api/request"
	"belajar-api/response"

	"github.com/gin-gonic/gin"
)

func GetRatings(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))
	keywordNam, keyword := helper.GetKeywordQueryFromRequest(c)
	orderName, orderBy := helper.GetOrderQueryFromRequest(c)
	pagination := helper.GeneratePagination(c)

	ratings, err := models.GetAllRatings(id, keywordNam, keyword, orderName, orderBy, pagination)

	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.RatingResponseList(c, "Rating Data Successfully Displayed", ratings, orderName, orderBy, pagination)
}

func CreateBookRating(c *gin.Context) {
	var BookRatingRequest request.BookRatingRequest
	err := c.ShouldBindJSON(&BookRatingRequest)
	if helper.ShouldBindJSONError(c, err) {
		return
	}

	newRating := models.Rating{
		BookID: BookRatingRequest.BookID,
		Rating: BookRatingRequest.Rating,
		Review: BookRatingRequest.Review,
	}

	r, err := models.CreateRating(newRating)

	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.RatingResponse1(c, "Successfully Created New Rating", r)
}

package controllers

import (
	"belajar-api/helper"
	"belajar-api/models"
	"belajar-api/request"
	"belajar-api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGenres(c *gin.Context) {
	keywordName, keyword := helper.GetKeywordQueryFromRequest(c)
	orderName, orderBy := helper.GetOrderQueryFromRequest(c)
	pagination := helper.GeneratePagination(c)

	genre, err := models.GetAllGenres(keywordName, keyword, []string{}, orderName, orderBy, pagination)

	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.GenreResponseList(c, "Genre Data Successfully Displayed", genre, orderName, orderBy, pagination)
}

func GetGenreByID(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))

	genre, err := models.GetGenreByID(id)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.GenreResponse1(c, "Genre Data Successfully Displayed", genre)
}

func CreateGenre(c *gin.Context) {
	var genreRequest request.GenreRequest

	err := c.ShouldBindJSON(&genreRequest)
	if err != nil {
		helper.ShouldBindJSONError(c, err)
		return
	}

	newGenre := models.Genre{
		Name:        genreRequest.Name,
		Description: genreRequest.Description,
	}

	g, err := models.CreateGenre(newGenre)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.GenreResponse1(c, "Genre Data Successfully Created", g)
}

func UpdateGenre(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))
	var genreRequest request.GenreRequest

	err := c.ShouldBindJSON(&genreRequest)
	if err != nil {
		helper.ShouldBindJSONError(c, err)
		return
	}

	updateGenre := models.Genre{
		ID:          id,
		Name:        genreRequest.Name,
		Description: genreRequest.Description,
	}

	g, err := models.UpdateGenre(updateGenre)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.GenreResponse1(c, "Genre Data Successfully Updated", g)
}

func DeleteGenre(c *gin.Context) {
	id, _ := helper.StringToInt(c.Param("id"))

	err := models.DeleteGenre(id)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	c.JSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Successfully Deleted Existing Genre Data",
	})
}

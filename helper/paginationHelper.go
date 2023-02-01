package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page  int
	Limit int
}

func GeneratePagination(c *gin.Context) Pagination {
	page := c.Request.URL.Query().Get("Page")
	limit := c.Request.URL.Query().Get("Limit")
	pageint, _ := strconv.Atoi(page)
	limitint, _ := strconv.Atoi(limit)

	return Pagination{
		Page:  pageint,
		Limit: limitint,
	}
}

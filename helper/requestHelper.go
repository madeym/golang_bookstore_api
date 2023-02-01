package helper

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetKeywordQueryFromRequest(c *gin.Context) ([]string, string) {
	keywordName := c.Request.URL.Query().Get("KeywordName")
	keyword := c.Request.URL.Query().Get("Keyword")

	temp := strings.Split(keywordName, ",")

	var columnName = make([]string, len(temp))

	for idx, value := range temp {
		columnName[idx] = value
	}

	return columnName, keyword
}

func GetParamQueryFromRequest(c *gin.Context, param []string, condition []string, value []string) []string {
	var result []string
	for idx, v := range param {
		temp := c.Request.URL.Query().Get(v)
		var val string
		if temp != "" {
			if value[idx] == "" {
				val = temp
			} else {
				val = value[idx]
			}
			if strings.Contains(condition[idx], "LIKE") {
				val = "%" + val + "%"
			}
			if value[idx] == "" {
				if strings.Contains(condition[idx], "LIKE") {

				}
				result = append(result, temp)
			} else {
				result = append(result, value[idx])
			}
			result = append(result, condition[idx])
		}
	}
	return result
}

func GetOrderQueryFromRequest(c *gin.Context) (string, string) {
	orderName := c.Request.URL.Query().Get("OrderName")
	orderBy := c.Request.URL.Query().Get("OrderBy")

	return orderName, orderBy
}

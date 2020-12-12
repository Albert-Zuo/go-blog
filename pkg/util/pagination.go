package util

import (
	"github.com/Albert/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

var PageQueryName = "page"

// GetPage get pagination DBPage of request query page
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query(PageQueryName)).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}

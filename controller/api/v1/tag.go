package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iScript/app/model"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	data["list"], _ = model.GetTags(1, 10, maps)
	data["total"], _ = model.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	model.AddTag("test", 9, "123123123123")
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}

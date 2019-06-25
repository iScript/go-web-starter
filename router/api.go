package router

import (
	v1 "../controller/api/v1"
	"../middleware/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("v1")
	{
		apiv1.GET("/index", v1.Index)

		//获取标签列表
		apiv1.GET("/tag", v1.AddTag)
		//新建标签
		apiv1.POST("/tag", v1.AddTag)
		//更新指定标签
		//apiv1.PUT("/tag/:id", v1.EditTag)
		//删除指定标签
		//apiv1.DELETE("/tag/:id", v1.DeleteTag)

		apiv1.GET("/login", v1.Login)

		user := apiv1.Group("user")
		user.Use(jwt.JWT())
		{
			user.GET("info", v1.GetUser)
		}

	}

	return r
}

package v1

import (
	"fmt"
	"net/http"

	"../../../model"
	"../../../pkg/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	data := make(map[string]interface{})

	isExist, err := model.UserAuth(username, password)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10041, "message": "login", "data": data})
		return
	}

	if !isExist {
		c.JSON(http.StatusOK, gin.H{"code": 10042, "message": "login", "data": data})
		return
	}

	fmt.Println("1111", username, password)

	token, err := util.GenerateToken(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 10043, "message": "login", "data": data})
	} else {
		data["token"] = token

		c.JSON(http.StatusOK, gin.H{"code": 10045, "message": "login", "data": data})
	}

}

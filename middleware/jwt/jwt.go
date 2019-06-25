package jwt

import (
	"fmt"
	"net/http"
	"time"

	"../../pkg/util"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 0

		token := c.Query("token")
		if token == "" {
			code = 4031
		} else {
			claims, err := util.ParseToken(token)

			fmt.Println(claims)

			if err != nil {
				code = 4032
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 4033
			}
		}

		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": "jwt fail",
				"data":    data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

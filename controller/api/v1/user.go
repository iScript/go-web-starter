package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "uuuu"})

}

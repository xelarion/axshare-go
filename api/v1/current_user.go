package v1

import (
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
)

func CurrentUserId(c *gin.Context) (userId uint) {
	tokenString := utils.GetHeaderToken(c)
	userId, err := utils.GetUserIdByToken(tokenString)

	if err != nil {
		panic(err)
	}

	return userId
}

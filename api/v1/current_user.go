package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
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

func CurrentAccount(c *gin.Context) (account models.Account) {
	db.AxshareDb.First(&account, CurrentUserId(c))
	return account
}

func CurrentUser(c *gin.Context) (user models.User) {
	db.AxshareDb.First(&user, CurrentUserId(c))
	return user
}

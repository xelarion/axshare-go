package v1

import (
	"axshare_go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
)

func CurrentAccountId(c *gin.Context) (userId uint) {
	tokenString := utils.GetHeaderToken(c)
	acct.Finder.FindAccountIdByToken(tokenString)
	return acct.Finder.FindAccountIdByToken(tokenString)
}

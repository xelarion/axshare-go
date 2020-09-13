package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xandercheung/acct"
	"github.com/xandercheung/ogs-go"
	"gorm.io/gorm"
	"strconv"
)

func ResetAccountPassword(g *gin.Context) {
	account, err := loadAccount(g)
	if err != nil {
		return
	}

	account.Password = "123456"
	if err = account.UpdatePassword(); err != nil {
		acct.Utils.JSON(g, ogs.RspError(ogs.StatusUpdateFailed, err.Error()))
	} else {
		acct.Utils.JSON(g, ogs.RspOK("Reset Password Successfully"))
	}
}

func loadAccount(g *gin.Context) (account acct.Account, err error) {
	id, _ := strconv.Atoi(g.Param("id"))
	account = acct.Finder.FindAccountById(uint(id))

	if !account.IsPersisted() {
		acct.Utils.JSON(g, ogs.RspError(ogs.StatusUserNotFound, "Account Not Found"))
		return account, gorm.ErrRecordNotFound
	}

	return account, nil
}

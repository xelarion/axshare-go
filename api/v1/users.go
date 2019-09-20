package v1

import (
	"axshare_go/internal/models"
	"axshare_go/internal/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
)

// 授权登录
func Login(c *gin.Context) {
	account := &models.Account{}
	err := json.NewDecoder(c.Request.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		c.JSON(http.StatusOK,
			ogs.RspBase(ogs.StatusSystemError, ogs.ErrorMessage("Invalid Request")))
		return
	}

	resp := models.Authenticate(account.Email, account.Username, account.Password)
	c.JSON(http.StatusOK, resp)
}

// 销毁授权
func Logout(c *gin.Context) {
	account := models.FindAccountByToken(utils.GetHeaderToken(c))
	err := account.DestroyToken()
	if err != nil {
		c.JSON(http.StatusOK, ogs.RspOK(ogs.ErrorMessage("操作失败！")))
	}
	c.JSON(http.StatusOK, ogs.RspOK(ogs.SuccessMessage("退出成功！")))
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	user := models.FindAccountByToken(utils.GetHeaderToken(c))
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), user))
}

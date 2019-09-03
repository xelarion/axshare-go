package user

import (
	"axshare_go/internal/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ogsapi/ogs-go"
	"net/http"
)

func Authenticate(c *gin.Context) {
	account := &models.Account{}
	err := json.NewDecoder(c.Request.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		c.JSON(http.StatusOK,
			ogs.RspBase(ogs.StatusInvalidToken, ogs.ErrorMessage("Invalid Token")))
		return
	}

	resp := models.Login(account.Email, account.Password)
	c.JSON(http.StatusOK, resp)
}

package v1

import (
	"axshare_go/internal/db"
	"axshare_go/internal/models"
	"axshare_go/internal/pg"
	"axshare_go/internal/utils"
	"encoding/json"
	"fmt"
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
	account := CurrentAccount(c)
	err := account.DestroyToken()
	if err != nil {
		c.JSON(http.StatusOK, ogs.RspOK(ogs.ErrorMessage("操作失败！")))
	}
	c.JSON(http.StatusOK, ogs.RspOK(ogs.SuccessMessage("退出成功！")))
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	user := CurrentUser(c)
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), FormatUser(user)))
}

func GetUsers(c *gin.Context) {
	var users []models.User
	relation := db.AxshareDb.Model(&models.User{})
	relation, paginate := pg.PaginateGin(relation, c)
	relation.Find(&users)

	c.JSON(http.StatusOK, ogs.RspOKWithPaginate(
		ogs.BlankMessage(), FormatUserList(users), paginate))
}

func GetUser(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	id := utils.ParseUint(c.Param("id"))
	user := models.User{}
	db.AxshareDb.First(&user, id)
	c.JSON(http.StatusOK, ogs.RspOKWithData(ogs.BlankMessage(), FormatUser(user)))
}

func UpdateUser(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	id := utils.ParseUint(c.Param("id"))
	account := models.Account{}
	db.AxshareDb.First(&account, id)

	params := utils.GetBodyParams(c)
	if params["password"] != nil {
		params["encrypted_password"] = utils.PasswordToBcrypt(params["password"].(string))
	}
	db.AxshareDb.Model(&account).Updates(params)

	c.JSON(http.StatusOK, ogs.RspOK(ogs.BlankMessage()))
}

func CreateUser(c *gin.Context) {
	params := utils.GetBodyParams(c)

	account := models.Account{}
	jsonBody, _ := json.Marshal(params)
	_ = json.Unmarshal(jsonBody, &account)
	fmt.Println(params)
	account.Create()

	c.JSON(http.StatusOK, ogs.RspOK(ogs.SuccessMessage("创建成功！")))
}

func FormatUserList(users []models.User) []map[string]interface{} {
	var json = make([]map[string]interface{}, len(users))
	for i, user := range users {
		data := FormatUser(user)
		json[i] = data
	}
	return json
}

func FormatUser(user models.User) map[string]interface{} {
	var data = make(map[string]interface{})
	data["id"] = user.ID
	data["nickname"] = user.Nickname
	data["username"] = user.Username
	data["email"] = user.Email
	data["avatar"] = user.Avatar
	data["created_at"] = utils.FormatDateTime(user.CreatedAt)
	data["updated_at"] = utils.FormatDateTime(user.UpdatedAt)
	return data
}

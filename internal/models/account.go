package models

import (
	"axshare_go/internal/db"
	"axshare_go/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/ogsapi/ogs-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

//a struct to rep user account
type Account struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"json:"email"`
	Username string `gorm:"type:varchar(100);unique_index"json:"username"`
	Password string `gorm:"column:encrypted_password" json:"password"`
	Token    string `gorm:"-" json:"token"`
}

func (Account) TableName() string {
	return "users"
}

func Authenticate(email, username, password string) interface{} {
	account := Account{}
	err := db.AxshareDb.Where(&Account{Email: email, Username: username}).First(&account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ogs.RspBase(ogs.StatusUserNotFound, ogs.ErrorMessage("用户不存在！"))
		}
		return ogs.RspBase(ogs.StatusSystemError, ogs.ErrorMessage("Connection error. Please retry"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword { // Password does not match!
			return ogs.RspBase(ogs.StatusErrorPassword, ogs.ErrorMessage("密码错误！"))
		}
		return ogs.RspBase(ogs.StatusSystemError, ogs.ErrorMessage(err.Error()))
	}

	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	account.GenToken()

	return ogs.RspOKWithData(ogs.SuccessMessage("登录成功！"), account)
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return utils.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := db.AxshareDb.Model(&Account{}).Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}

	return utils.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	db.AxshareDb.Create(account)

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	account.GenToken()

	account.Password = "" //delete password

	response := utils.Message(true, "Account has been created")
	response["account"] = account
	return response
}

// 生成 token
func (account *Account) GenToken() string {
	tk := &utils.TokenClaims{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_KEY")))
	account.Token = tokenString
	return tokenString
}

// todo
func (account *Account) DestroyToken() error {
	return nil
}


package models

import (
	"axshare_go/internal/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/ogsapi/ogs-go"
	"golang.org/x/crypto/bcrypt"
	"os"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	ID       uint   `gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

func (Account) TableName() string {
	return "users"
}

func Login(email, password string) interface{} {

	account := Account{}
	err := db.AxshareDb.Debug().Where("email = ?", email).First(&account).Error
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
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	return ogs.RspOKWithData(ogs.SuccessMessage("登录成功！"), account)
}
